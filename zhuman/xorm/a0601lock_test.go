package main

import (
	"log"
	"testing"
)

func TestLockPrepare(t *testing.T) {
	orders := make([]*UserLock, 2)
	orders[0] = new(UserLock)
	orders[0].Name = "name0"

	orders[1] = new(UserLock)
	orders[1].Name = "aaa"
	affected, err := engine.Insert(&orders)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

// 乐观锁
func TestLock(t *testing.T) {
	var user UserLock
	engine.ID(1).Get(&user)
	user.Name = user.Name + "_1"
	// SELECT * FROM user WHERE id = ?
	engine.ID(1).Update(&user)
	// UPDATE user SET ..., version = version + 1 WHERE id = ? AND version = ?
}

/*
2022-07-03T13:17:18.567730Z       131 Close stmt
2022-07-03T13:17:18.575643Z       131 Prepare   SELECT `id`, `name`, `version` FROM `user_lock` WHERE `id`=? LIMIT 1
2022-07-03T13:17:18.635755Z       131 Execute   SELECT `id`, `name`, `version` FROM `user_lock` WHERE `id`=1 LIMIT 1
2022-07-03T13:17:18.707647Z       131 Close stmt
2022-07-03T13:17:18.708043Z       131 Prepare   UPDATE `user_lock` SET `name` = ?, `version` = `version` + 1 WHERE `id`=? AND `version`=?
2022-07-03T13:17:18.771435Z       131 Execute   UPDATE `user_lock` SET `name` = 'name0_1', `version` = `version` + 1 WHERE `id`=1 AND `version`=1
*/
