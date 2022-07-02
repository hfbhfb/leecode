package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 业务1 抢购 `使用mysql接口+redis配合`
// Type1 `不使用分布式锁` + `只将数据库数据同步到redis`
func TestQianggouMysqlPlusRedisType1(t *testing.T) {

	print("main here")
	chMysql = make(chan int, 99)
	chRedis = make(chan int, 999)
	countMysqlPrepare()

	iSecondsTotal := 10
	iSecondsTest := iSecondsTotal
	t1 := time.NewTicker(1 * time.Second)

	// 并发数
	Concurrence := 3
	for i := 0; i < Concurrence; i++ {
		go func() {
			for {
				ProductInfo := simuRedisOptRead(true)
				throughPutRate++

				if ProductInfo == StatusProNotExist {
					simuMysqlOptOnlyConsummer()
					simuRedisOptWriteFakelInfo(allPocucts)

				} else if ProductInfo == StatusProZero {
				} else {

				}

				// if flagHaveProduct {
				// 	productIdGeted := simuMysqlOpt(false)
				// 	mu.Lock()
				// 	if productIdGeted >= 0 {
				// 		countGetProduct++
				// 	}
				// 	mu.Unlock()
				// }
			}
		}()
	}
ForEnd:
	for {
		select {
		case <-t1.C:
			fmt.Println("吞吐总数:", throughPutRate)
			fmt.Println("吞吐率为:", throughPutRate/(iSecondsTotal+1-iSecondsTest))
			fmt.Println("并发数为:", Concurrence)
			fmt.Println("抢购到的商品数:", countGetProduct)
			fmt.Println("mysql触发率:", allMysqltrick)
			fmt.Println("redis触发率:", allRedistrick)

			iSecondsTest--
			if iSecondsTest <= 0 {
				break ForEnd
			}
			fmt.Println("程序在 ", iSecondsTest, "秒 后结束")
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("吞吐总数:", throughPutRate)
	fmt.Println("吞吐率为tps:", throughPutRate/(iSecondsTotal+1-iSecondsTest))
	fmt.Println("并发数为:", Concurrence)
	fmt.Println("抢购到的商品数:", countGetProduct)
	fmt.Println("")
	fmt.Println("商品的发放总数:", allPocucts)
	fmt.Println("mysql触发率:", allMysqltrick)
	fmt.Println("mysql触发率 和 吞吐总数不一致是正常的 , 因为 100毫秒触发一次mysql/redis处理能力")

	fmt.Println("程序结束")

}

// 业务1 抢购 `只使用mysql接口`
func TestQianggouOnlyMysql(t *testing.T) {

	print("main here")
	chMysql = make(chan int, 99)
	chRedis = make(chan int, 999)
	countMysqlPrepare()

	iSecondsTotal := 10
	iSecondsTest := iSecondsTotal
	t1 := time.NewTicker(1 * time.Second)

	// 并发数
	Concurrence := 3
	for i := 0; i < Concurrence; i++ {
		go func() {
			for {
				productIdGeted := simuMysqlOpt(false)
				mu.Lock()
				throughPutRate++
				if productIdGeted >= 0 {
					countGetProduct++
				}
				mu.Unlock()
			}
		}()
	}
ForEnd:
	for {
		select {
		case <-t1.C:
			fmt.Println("吞吐总数:", throughPutRate)
			fmt.Println("吞吐率为:", throughPutRate/(iSecondsTotal+1-iSecondsTest))
			fmt.Println("并发数为:", Concurrence)
			fmt.Println("抢购到的商品数:", countGetProduct)
			fmt.Println("mysql触发率:", allMysqltrick)

			iSecondsTest--
			if iSecondsTest <= 0 {
				break ForEnd
			}
			fmt.Println("程序在 ", iSecondsTest, "秒 后结束")
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("吞吐总数:", throughPutRate)
	fmt.Println("吞吐率为tps:", throughPutRate/(iSecondsTotal+1-iSecondsTest))
	fmt.Println("并发数为:", Concurrence)
	fmt.Println("抢购到的商品数:", countGetProduct)
	fmt.Println("")
	fmt.Println("商品的发放总数:", allPocucts)
	fmt.Println("mysql触发率:", allMysqltrick)
	fmt.Println("mysql触发率 和 吞吐总数不一致是正常的 , 因为 100毫秒触发一次mysql/redis处理能力")

	fmt.Println("程序结束")

}

var (
	allPoducts = 0 // 抢购未开始,商品数为0
	mu         sync.Mutex

	// 模拟mysq秒级50个接口请求
	MaxSqlWrite = 50
	chMysql     chan int

	throughPutRate  = 0 // 吞吐率测试
	countGetProduct = 0 // 抢购到的商品

	allMysqltrick = 0 // mysql触发总数
	allRedistrick = 0 // redis触发总数
	allPocucts    = 0 // 一共放了多少件商品到抢购

	productInfoInt = -1 // 默认redis没有值

	// 模拟redis秒级2000个处理请求
	MaxRedis = 2000
	chRedis  chan int
)

func countMysqlPrepare() {
	// 模拟mysql每秒50个请求处理能力
	go func() {
		t1 := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t1.C:
				// fmt.Println(len(chMysql))
				if len(chMysql) >= MaxSqlWrite {
				} else {
					for i := 0; i < int((MaxSqlWrite / 10)); i++ {
						allMysqltrick++
						chMysql <- 1
					}
				}

				if len(chRedis) >= MaxSqlWrite {
				} else {
					for i := 0; i < int((MaxRedis / 10)); i++ {
						allRedistrick++
						chRedis <- 1
					}
				}
			}
		}
	}()

	// 模拟每3秒 投放5个商品到篮子里,蓝子里最多5个商品
	go func() {
		t1 := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-t1.C:
				mu.Lock()
				for i := 0; i < 5; i++ {
					if allPoducts < 5 {
						allPocucts++
						allPoducts++
					}
				}
				mu.Unlock()
			}
		}
	}()
}

// longConn 是否长连接
func simuMysqlOpt(longConn bool) int {

	if longConn == false {
		// 模拟长连接短连接区别
		time.Sleep(20 * time.Millisecond)
	}
	<-chMysql

	mu.Lock()
	productId := -1
	if allPoducts > 0 {
		allPoducts--
		// 随机一个产品id
		productId = 9
	}
	mu.Unlock()
	return productId
}

func simuMysqlOptOnlyConsummer() {
	<-chMysql
}

var (
	StatusProNotExist = -1
	StatusProZero     = 0
)

// 三态: -1: redis数据库没有对应的值[要去读一下数据库]   0: 没有产品   1-n: 里面有产品
func simuRedisOptRead(longConn bool) int {
	if longConn == false {
		// 模拟长连接短连接区别
		time.Sleep(20 * time.Millisecond)
	}
	<-chRedis

	mu.Lock()

	if allPoducts > 0 {
		mu.Unlock()
		return productInfoInt
	}
	mu.Unlock()
	return productInfoInt
}

func simuRedisOptClearFakeInfo() bool {
	mu.Lock()
	productInfoInt = -1
	mu.Unlock()
	return true
}

func simuRedisOptWriteFakelInfo(products int) bool {
	mu.Lock()
	productInfoInt = products
	mu.Unlock()
	return true
}
