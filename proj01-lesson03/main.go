package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

// percentTarget 目标占用百分比
func cpucontrol(percentTarget float64) {

	done := make(chan int)

	doLoopNum := runtime.NumCPU()
	signerPercent := (1 / float64(doLoopNum)) * 100
	fmt.Println(signerPercent)
	doLoopNum = int(percentTarget/signerPercent) + 1 // 目标几核
	fmt.Println("doLoopNum:", doLoopNum)

	for i := 0; i < doLoopNum; i++ {
		tmpi := i
		go func() {
			if tmpi == (doLoopNum - 1) {
				it := 0
				tmpt := int((float64((int(percentTarget) % int(signerPercent))) / signerPercent) * 10)
				fmt.Println("tmpt:", tmpt)
				fmt.Println("a:", (int(percentTarget) % int(signerPercent)))
				fmt.Println("b:", (float64((int(percentTarget) % int(signerPercent))) / signerPercent))
				fmt.Println("c:", (float64((int(percentTarget)%int(signerPercent)))/signerPercent)*10)

				t1 := time.NewTicker(100 * time.Millisecond)
				for {
					select {
					case <-done:
						return
					case <-t1.C:
						it++
						if it < tmpt {
							time.Sleep(time.Millisecond * 100)
							it++
						}
						it = it % 10
						// 0.1 秒判断一下最后一个协程需要不需要休息一下,做到精准cpu控制

					default:
					}
				}
			} else {
				for {
					select {
					case <-done:
						return
					default:
					}
				}
			}

		}()
	}

	time.Sleep(time.Second * 30)
	close(done)
}

func memcontrol(percentTarget float64) {
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Getpagesize())
}

func indexHandlerHelloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func httpopt() {
	fs := http.FileServer(http.Dir("./")) //去静态目录找 得到fs对象：文件服务器
	http.HandleFunc("/helloworld", indexHandlerHelloworld)
	http.Handle("/", http.StripPrefix("/", fs)) //过滤掉/static/这部分，把剩下的部分给fs
	http.ListenAndServe(":8099", nil)
}

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	httpopt()
	// memcontrol(70)
	// cpucontrol(1)
}
