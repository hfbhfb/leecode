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
