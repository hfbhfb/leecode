package main

import (
	"log"
	"testing"
)

func TestTransationLock2(t *testing.T) {

	err := func() error {
		// var err error

		session := engine.NewSession()
		defer session.Close()

		// add Begin() before any action
		if err := session.Begin(); err != nil {
			return err
		}

		userslice := make([]UserinfoTrans, 0)
		// session.Where("id = ?", 4).Find(&userslice)
		session.Where("id = ?", 4).ForUpdate().Find(&userslice)
		log.Println(userslice)
		// add Commit() after all actions
		return session.Commit()
	}()
	if err != nil {
		log.Println(err.Error())
		t.Error(err.Error())
	}
}

/*
2022-07-04T23:18:38.726987Z       182 Query     START TRANSACTION
2022-07-04T23:18:38.824951Z       182 Prepare   SELECT `id`, `username`, `departname`, `alias`, `created` FROM `userinfo_trans` WHERE (id = ?) FOR UPDATE
2022-07-04T23:18:38.944656Z       182 Execute   SELECT `id`, `username`, `departname`, `alias`, `created` FROM `userinfo_trans` WHERE (id = 4) FOR UPDATE
2022-07-04T23:18:39.027283Z       182 Close stmt
2022-07-04T23:18:39.027534Z       182 Query     COMMIT
*/
