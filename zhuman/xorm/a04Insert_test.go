package main

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	user := new(User)
	user.Name = "myname11"
	affected, err := engine.Insert(user)
	// INSERT INTO user (name) values (?)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

// 2022-07-03T08:32:57.693781Z        63 Prepare   INSERT INTO `user` (`usr_name`,`email`) VALUES (?,?)
// 2022-07-03T08:32:57.745467Z        63 Execute   INSERT INTO `user` (`usr_name`,`email`) VALUES ('myname','')
// 2022-07-03T08:32:57.845901Z        63 Close stmt

func TestInsertSlice(t *testing.T) {
	users := make([]*User, 2)
	users[0] = new(User)
	users[0].Name = "name0"

	users[1] = new(User)
	users[1].Name = "name2"
	affected, err := engine.Insert(&users)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

// 2022-07-03T08:35:56.075942Z        66 Prepare   INSERT INTO `user` (`usr_name`,`email`) VALUES (?, ?),(?, ?)
// 2022-07-03T08:35:56.155462Z        66 Execute   INSERT INTO `user` (`usr_name`,`email`) VALUES ('name0', ''),('name2', '')

/*
// Go语言规范不允许这样做，因为两种类型在内存中没有相同的表现形式。
func TestInsertMultiLines(t *testing.T) {
	users := make([]*User, 2)
	users[0] = new(User)
	users[0].Name = "name0"

	users[1] = new(User)
	users[1].Name = "name2"
	affected, err := engine.Insert(users...)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}
*/

func TestInsertMultiItem(t *testing.T) {
	users := make([]User, 1)
	users[0].Name = "name9"
	questions := make([]Question, 1)
	questions[0].Content = "whywhywhwy?"
	affected, err := engine.Insert(&users, &questions)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

/*
user失败后
2022-07-03T09:11:17.519340Z        71 Close stmt
2022-07-03T09:11:17.527532Z        71 Prepare   INSERT INTO `user` (`usr_name`,`email`) VALUES (?, ?)
2022-07-03T09:11:17.580480Z        71 Execute   INSERT INTO `user` (`usr_name`,`email`) VALUES ('name7', '')
2022-07-03T09:11:17.679032Z        71 Close stmt
*/
/*
2022-07-03T09:19:31.755203Z        72 Prepare   INSERT INTO `user` (`usr_name`,`email`) VALUES (?, ?)
2022-07-03T09:19:31.835060Z        72 Execute   INSERT INTO `user` (`usr_name`,`email`) VALUES ('name9', '')
2022-07-03T09:19:31.946759Z        72 Close stmt
2022-07-03T09:19:31.954916Z        72 Prepare   INSERT INTO `question` (`content`) VALUES (?)
2022-07-03T09:19:32.026725Z        72 Execute   INSERT INTO `question` (`content`) VALUES ('whywhywhwy?')
*/

func TestInsertAutoCreate(t *testing.T) {
	users := make([]TAutoCreated, 1)
	affected, err := engine.Insert(&users)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}
