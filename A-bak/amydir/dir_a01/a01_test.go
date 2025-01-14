package dira1

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// count 数组rune 字符出现次数
func TestA01(t *testing.T) {

	mmStr := []string{"abc", "c", "中文", "文"}

	start := time.Now() // 获取当前时间
	elapsed := time.Now().Sub(start)

	more := 90080000
	mmStr = make([]string, more)
	for i := 0; i < more; i++ {
		mmStr[i] = "abc"
	}

	/*
		for i := 0; i < more; i++ {
			mmStr = append(mmStr, "abc")
		}
	*/
	elapsed = time.Now().Sub(start)
	fmt.Println("申请空间:", elapsed)

	start = time.Now()        // 获取当前时间
	aa := CountLetters(mmStr) // 9千万条时候,5秒
	elapsed = time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)

	start = time.Now()                // 获取当前时间
	bb := CountLettersCurrency(mmStr) // 9千万条时候,1秒
	elapsed = time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时 Currency：", elapsed)
	// t.Log(aa)
	t.Log(aa == nil)
	// t.Log(bb)
	t.Log(bb == nil)

}

type LetterFre map[rune]int

func CountLetters(strs []string) LetterFre {

	ret := make(LetterFre)
	for _, str := range strs {
		for _, r := range str {
			ret[r]++
		}

	}
	return ret
}

// count 数组rune 字符出现次数,并发
func CountLettersCurrency(strs []string) LetterFre {

	wg := sync.WaitGroup{}

	fmt.Println("字符串长度:", len(strs))
	// fmt.Println(cap(strs))
	cpus := runtime.NumCPU()

	retAll := make(LetterFre)
	var mu sync.Mutex

	iDeal := 0
	for i := 0; i < cpus; i++ {

		len1 := (len(strs) + cpus) / cpus
		// fmt.Println(len1)
		ci := make(chan []string, 200)
		fromInt := i * len1
		endInt := (i + 1) * len1
		if endInt > len(strs) {
			endInt = len(strs)
		}

		if fromInt > len(strs) {
		} else {
			ci <- strs[fromInt:endInt]
		}

		fmt.Println(fromInt)
		fmt.Println(endInt)

		close(ci)

		// fmt.Println(i)
		wg.Add(1)
		go func(cstrs <-chan []string) {

			ret := make(LetterFre)
		LOOP:
			for {
				select {
				case strs, err := <-cstrs:
					if err != true {
						// fmt.Println("break")
						break LOOP
					}
					for _, str := range strs {
						for _, r := range str {

							// 如果在这里操作,会30秒超时
							// mu.Lock()
							// for k, v := range ret {
							// 	retAll[k] = v
							// }
							// mu.Unlock()

							ret[r]++
						}

					}

					iDeal += len(strs)
					// fmt.Println(len(strs))
				}
			}

			mu.Lock()
			for k, v := range ret {
				retAll[k] = v
			}
			mu.Unlock()

			wg.Done()
		}(ci)
	}

	wg.Wait()
	// fmt.Println(cpus)
	// fmt.Println(iDeal)
	fmt.Println(retAll)
	// ret := make(LetterFre)
	return nil
}
