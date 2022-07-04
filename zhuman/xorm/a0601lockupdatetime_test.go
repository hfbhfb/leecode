package main

import (
	"testing"
)

// 乐观锁
func TestLockUpdateTime(t *testing.T) {
	var user UserLock
	engine.ID(1).Get(&user)
	user.Name = user.Name + "_1"
	// SELECT * FROM user WHERE id = ?
	engine.ID(1).Update(&user)
	// UPDATE user SET ..., version = version + 1 WHERE id = ? AND version = ?
}

/*
2022-07-03T13:22:21.765500Z       133 Prepare   SELECT `id`, `name`, `version`, `updated_at` FROM `user_lock` WHERE `id`=? LIMIT 1
2022-07-03T13:22:21.817323Z       133 Execute   SELECT `id`, `name`, `version`, `updated_at` FROM `user_lock` WHERE `id`=1 LIMIT 1
2022-07-03T13:22:21.877300Z       133 Close stmt
2022-07-03T13:22:21.885717Z       133 Prepare   UPDATE `user_lock` SET `name` = ?, `updated_at` = ?, `version` = `version` + 1 WHERE `id`=? AND `version`=?
2022-07-03T13:22:21.937307Z       133 Execute   UPDATE `user_lock` SET `name` = 'name0_1_1', `updated_at` = '2022-07-03 21:22:21', `version` = `version` + 1 WHERE `id`=1 AND `version`=2
*/
