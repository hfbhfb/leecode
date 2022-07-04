package main

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPrepareAlias(t *testing.T) {
	orders := make([]*Order, 2)
	orders[0] = new(Order)
	orders[0].Name = "name0"

	orders[1] = new(Order)
	orders[1].Name = "aaa"
	affected, err := engine.Insert(&orders)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

// 给Table设定一个别名
func TestAlias(t *testing.T) {
	order := new(Order)
	name := "aaa"
	flag, err := engine.Alias("o").Where("o.name = ?", name).Get(order)
	// INSERT INTO user (name) values (?)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", flag)
	if flag {
		log.Println("order info : ", order)
	}
}

/*
2022-07-03T09:46:20.606720Z        80 Prepare   SELECT `id`, `name`, `time`, `created_at` FROM `order` AS `o` WHERE (o.name = ?) LIMIT 1
2022-07-03T09:46:20.658330Z        80 Execute   SELECT `id`, `name`, `time`, `created_at` FROM `order` AS `o` WHERE (o.name = 'aaa') LIMIT 1
*/

// 条件 And(string, …interface{})

// 给Table设定一个别名
func TestAnd(t *testing.T) {
	order := new(Order)
	name := "aaa"
	flag, err := engine.Alias("o").Where("o.name = ?", name).And("o.id = ?", 2).Get(order)
	// INSERT INTO user (name) values (?)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", flag)
	if flag {
		log.Println("order info : ", order)
	}
}

/*
2022-07-03T09:51:19.648404Z        83 Prepare   SELECT `id`, `name`, `time`, `created_at` FROM `order` AS `o` WHERE (o.name = ?) AND (o.id = ?) LIMIT 1
2022-07-03T09:51:19.700049Z        83 Execute   SELECT `id`, `name`, `time`, `created_at` FROM `order` AS `o` WHERE (o.name = 'aaa') AND (o.id = 2) LIMIT 1
*/

// Asc(…string) 指定字段名正序排序，可以组合
func TestAsc(t *testing.T) {
	orders := make([]Order, 0)
	log.Println("order start len : ", len(orders))
	err := engine.Asc("id").Find(&orders)
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Println("order info : ", orders)
	}
}

/*
2022-07-03T09:57:37.466762Z        88 Query     SELECT `id`, `name`, `time`, `created_at` FROM `order` ORDER BY `id` ASC
*/

// Desc(…string) 指定字段名逆序排序，可以组合
func TestDsc(t *testing.T) {
	orders := make([]Order, 0)
	log.Println("order start len : ", len(orders))
	err := engine.Desc("id").Find(&orders)
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Println("order info : ", orders)
	}
}

/*
2022-07-03T10:06:07.038345Z        99 Query     SELECT `id`, `name`, `time`, `created_at` FROM `order` ORDER BY `id` DESC
*/

// * ID(interface{}) 传入一个主键字段的值，作为查询条件，如
func TestIDOpt(t *testing.T) {
	var user User
	flag, err := engine.ID(100).Get(&user)
	if err != nil {
		t.Error(err.Error())
	}
	if flag {
		log.Println("user info : ", user)
	} else {
		log.Println("user info (error): ", user)

	}
}
