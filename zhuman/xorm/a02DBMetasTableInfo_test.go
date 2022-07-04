package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestPrintAllTableinfo(t *testing.T) {
	ts, _ := engine.DBMetas()

	for _, v := range ts {
		log.Println(v.Name)
		t, _ := engine.TableInfo(v)
		// log.Println("t:%v", t)
		fmt.Printf("t: %v\n", t)
	}

}

func TestDump(t *testing.T) {
	engine.DumpAll(os.Stdout)
}

func TestDumpTofile(t *testing.T) {

	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fileName := fmt.Sprintf("./testtmp/dump%v.sql", currentTime.Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	engine.DumpAll(file)
}
