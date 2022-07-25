package main

import "testing"

func TestDeleted(t *testing.T) {
	var user User
	engine.ID(5).Unscoped().Get(&user)
	// 此时将可以获得记录
	engine.ID(5).Unscoped().Delete(&user)
	// 此时将可以真正的删除记录
}

// 软删除Deleted
func TestDeletedSoft(t *testing.T) {
	var user User
	engine.ID(6).Get(&user)
	// 此时将可以获得记录
	engine.ID(6).Delete(&user)
	// 此时将可以真正的删除记录
}

func TestDeleteWhere(t *testing.T) {
	tmp := new(Order)
	engine.Where("name = ?", "aaa").Delete(tmp)
}

/*
2022-07-18T06:38:23.402347Z       219 Close stmt
2022-07-18T06:38:23.408759Z       219 Prepare   DELETE FROM `order` WHERE (name = ?)
2022-07-18T06:38:23.459364Z       219 Execute   DELETE FROM `order` WHERE (name = 'aaa')
2022-07-18T06:38:23.562336Z       219 Close stmt
*/
