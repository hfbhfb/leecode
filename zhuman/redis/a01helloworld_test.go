package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
)

var Conn *redis.Client

func init() {
	log.Println("redis init Here")
	Conn = redis.NewClient(&redis.Options{
		Addr:     "render.tpddns.cn" + ":" + "16016",
		Password: "yourpass123",
		DB:       5,
	})
	if _, err := Conn.Ping().Result(); err != nil {
		log.Println("Connect to redis client failed, err: ", err)
		// return
	}
	// return conn
}

type User struct {
	Name string
	Age  int
}

// func (p *User) String() string {
// 	return "user aaa"
// }

func TestHelloworld(t *testing.T) {

	// 不会被库识别
	Conn.Set("aabb", &User{"hfb", 18}, 1000000*time.Second)

	Conn.Set("aaccc", "111", 1000000*time.Second)
	Conn.Set("aadd", "111", -1)

	Conn.Set("aaee", fmt.Sprintf("%v", User{"hfb", 18}), -1)

}

/*
1656974970.663389 [0 14.26.13.55:56092] "auth" "(redacted)"
1656974970.663424 [5 14.26.13.55:56092] "select" "5"
1656974970.723418 [5 14.26.13.55:56092] "set" "aaccc" "111" "ex" "1000000"
1656974970.783012 [5 14.26.13.55:56092] "set" "aadd" "111"
*/
