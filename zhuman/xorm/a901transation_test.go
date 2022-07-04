package main

import (
	"log"
	"testing"
	"time"
)

func TestTransation(t *testing.T) {
	log.Println("aaaa start")
	time.Sleep(time.Second * 3)
	log.Println("aaaa start")
	time.Sleep(time.Second * 3)
	log.Println("aaaa start")
	time.Sleep(time.Second * 3)

	err := func() error {
		var err error

		session := engine.NewSession()
		defer session.Close()

		// add Begin() before any action
		if err := session.Begin(); err != nil {
			return err
		}

		user1 := UserinfoTrans{Username: "xiaoxiao", Departname: "dev", Alias: "lunny", Created: time.Now()}
		if _, err = session.Insert(&user1); err != nil {
			return err
		}
		user2 := UserinfoTrans{Username: "yyy"}
		if _, err = session.Where("id = ?", 2).Update(&user2); err != nil {
			return err
		}

		if _, err = session.Exec("delete from userinfo_trans where username = ?", user2.Username); err != nil {
			return err
		}

		// add Commit() after all actions
		return session.Commit()
	}()
	if err != nil {
		log.Println(err.Error())
		t.Error(err.Error())
	}
}

/*
2022-07-03T14:06:00.478813Z       162 Query     START TRANSACTION
2022-07-03T14:06:00.546962Z       162 Prepare   INSERT INTO `userinfo_trans` (`username`,`departname`,`alias`,`created`) VALUES (?,?,?,?)
2022-07-03T14:06:00.619060Z       162 Execute   INSERT INTO `userinfo_trans` (`username`,`departname`,`alias`,`created`) VALUES ('xiaoxiao','dev','lunny','2022-07-03 22:06:00')
2022-07-03T14:06:00.678871Z       162 Close stmt
2022-07-03T14:06:00.698924Z       162 Prepare   UPDATE `userinfo_trans` SET `username` = ? WHERE (id = ?)
2022-07-03T14:06:00.766836Z       162 Execute   UPDATE `userinfo_trans` SET `username` = 'yyy' WHERE (id = 2)
2022-07-03T14:06:00.847792Z       162 Close stmt
2022-07-03T14:06:00.855003Z       162 Prepare   delete from userinfo_trans where username = ?
2022-07-03T14:06:00.926730Z       162 Execute   delete from userinfo_trans where username = 'yyy'
2022-07-03T14:06:00.998795Z       162 Close stmt
2022-07-03T14:06:01.006953Z       162 Query     COMMIT
*/
