package main

import (
	"fmt"
	"log"
	"testing"
)

func TestRawsql(t *testing.T) {
	sql := "select * from user;"
	results, err := engine.Query(sql)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	for k, v := range results {
		fmt.Println(k, string(v["usr_name"]))
		// fmt.Println(k, v)
	}
}

/*
2022-07-03T13:45:30.340914Z       146 Query     select * from user
*/

func TestRawsqlExec(t *testing.T) {
	sql := "update `user` set usr_name=? where id=?"
	res, err := engine.Exec(sql, "name7_1", 6)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	log.Println(res.RowsAffected())
}

/*
2022-07-03T13:52:23.202326Z       149 Prepare   update `user` set usr_name=? where id=?
2022-07-03T13:52:23.254159Z       149 Execute   update `user` set usr_name='name7_1' where id=6
*/
