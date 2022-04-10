package dira04

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type LoadBalance struct {
	client []*Client
	size   int
}

func NewLoadBalance(size int) *LoadBalance {
	clients := make([]Client, size)
	loadBalance := &LoadBalance{
		client: make([]*Client, size),
		size:   size,
	}
	for i := 0; i < size; i++ {
		clients[i].name = fmt.Sprintf("Myname%v", i)

		loadBalance.client[i] = &clients[i]
	}
	return loadBalance
}
func (l *LoadBalance) getClient() *Client {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(l.size)
	fmt.Println(x)
	return l.client[x]
}

func (c *Client) Do() {
	// c.name = "kkk"
	// fmt.Println("do")
	fmt.Println(c.name)
}

type Client struct {
	name string
}

// LoadBalance 模型
func TestLxx(t *testing.T) {
	lb := NewLoadBalance(800)
	c := lb.getClient()
	c2 := lb.getClient()
	c.Do()
	c2.Do()
}
