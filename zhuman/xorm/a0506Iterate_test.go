package main

import (
	"log"
	"testing"
)

func TestIteratePrepare(t *testing.T) {
	orders := make([]*Userinfo, 2)
	orders[0] = new(Userinfo)
	orders[0].Name = "xlw"

	orders[1] = new(Userinfo)
	orders[1].Name = "xlw"
	affected, err := engine.Insert(&orders)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("Insert affected: ", affected)
}

func TestIterateHello(t *testing.T) {
	err := engine.Where("id > ? or name=?", 0, "xlw").Iterate(new(Userinfo), func(i int, bean interface{}) error {
		user := bean.(*Userinfo)
		//do somthing use i and user
		log.Println(user)
		return nil
	})
	if err != nil {
		log.Println(err.Error())
	}
}

/*
2022-07-03T10:20:58.333549Z       111 Prepare   SELECT `id`, `name`, `detail_id` FROM `userinfo` WHERE (id > ? or name=?)
2022-07-03T10:20:58.392291Z       111 Execute   SELECT `id`, `name`, `detail_id` FROM `userinfo` WHERE (id > 0 or name='xlw')
*/

func TestCount(t *testing.T) {
	user := new(User)
	total, err := engine.Where("id >?", 1).Count(user)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(total)
}

/*
2022-07-03T10:31:43.429825Z       112 Prepare   SELECT count(*) FROM `user` WHERE (id >?)
2022-07-03T10:31:43.543442Z       112 Execute   SELECT count(*) FROM `user` WHERE (id >1)
*/
