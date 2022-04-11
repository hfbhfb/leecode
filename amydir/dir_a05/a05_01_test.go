package dira05

import (
	"fmt"
	"strings"
	"testing"
)

// 外挂 路径 递归 寻址

type MapArrive struct {

	// 计算结果
	arrInfo           [][]bool
	arrInfoDirectPath [][]string

	// 原始数据
	minfo map[string][]string
	pos   map[string]string
}

func NewMapArrive(minfo map[string][]string, pos map[string]string) *MapArrive {
	a := &MapArrive{
		arrInfo: make([][]bool, 0),
		minfo:   minfo,
		pos:     pos,
	}
	namemap := make(map[string]int)
	countI := 0
	for k, _ := range minfo {
		a.arrInfo = append(a.arrInfo, make([]bool, len(minfo)))
		a.arrInfoDirectPath = append(a.arrInfoDirectPath, make([]string, len(minfo)))

		namemap[k] = countI
		countI++
	}
	for k, v := range pos {
		temA := strings.Split(k, "_")
		temB := strings.Split(v, "_")
		a.arrInfo[namemap[temA[0]]][namemap[temB[0]]] = true
		a.arrInfoDirectPath[namemap[temA[0]]][namemap[temB[0]]] = k + "->" + v
		fmt.Println(k)
		fmt.Println(v)
	}
	// for i := 0; i < len(minfo); i++ {
	// 	a.arrInfo = append(a.arrInfo, make([]bool, len(minfo)))
	// }

	for i := 0; i < len(minfo); i++ {
		a.arrInfo[i][i] = true
	}
	a.resolvePath()
	// for k, v := range minfo {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	fmt.Println(a)

	return a
}

func (p *MapArrive) resolvePath() {

	fmt.Println("------------")
	for i, v := range p.arrInfo {
		for i1, v1 := range v {
			if v1 == false {
				path := make([]int, len(p.arrInfo))
				for i, _ := range path {
					path[i] = -1
				}

				finish, res := p.resolvePoint(i, i1, path)
				if finish {
					fmt.Println("找到相应路径:", i+1, i1+1)
					_ = res
					for _, v := range res {
						fmt.Print(v+1, ",")
					}
					fmt.Println("")
				} else {
					fmt.Println("失败 未找到相应路径:", i+1, i1+1)
				}
			}
		}
	}
}

func (p *MapArrive) resolvePoint(from, to int, tmp []int) (bool, []int) {
	if checkLoop(from, tmp) {
		return false, nil
	}

	if p.arrInfo[from][to] {
		AddIndex(from, tmp)
		AddIndex(to, tmp)
		return true, tmp
	}
	AddIndex(from, tmp)

	for i, v := range p.arrInfo[from] {
		if v {

			path := make([]int, len(tmp))
			for i, _ := range tmp {
				path[i] = tmp[i]
			}

			finish, res := p.resolvePoint(i, to, path)
			if finish {
				return true, res
			}
		}
		// fmt.Println(v)
	}

	// fmt.Println(from, to)
	// fmt.Println(markfinishFind)
	// fmt.Println(markfinishFind)
	// fmt.Println(markfinishFind)
	return false, nil
}

func checkLoop(index int, mark []int) bool {
	for _, v := range mark {
		if v == index {
			return true
		}
	}
	return false
}

func AddIndex(index int, mark []int) {
	flag := false
	pos := -1
	for i, v := range mark {
		if v == index {
			flag = true
		}
		if v == -1 && pos == -1 {
			pos = i
		}
	}

	if flag == false {
		mark[pos] = index
		// *mark = append(*mark, index)
	}

}

func TestXxx(t *testing.T) {

	type Point struct {
		start string
		end   string
	}
	mapInfo := make(map[string][]string)
	mapInfo["map1"] = []string{
		"p1",
		"p2",
		"p3",
		"p4",
		"p5",
		"p6",
		"p7",
		"p8",
		"p9",
	}

	mapInfo["map2"] = []string{
		"p1",
		"p2",
		"p3",
		"p4",
		"p5",
		"p6",
		"p7",
		"p8",
		"p9",
	}

	mapInfo["map3"] = []string{
		"p1",
		"p2",
		"p3",
		"p4",
		"p5",
		"p6",
		"p7",
		"p8",
		"p9",
	}

	mapInfo["map4"] = []string{
		"p1",
		"p2",
		"p3",
		"p4",
		"p5",
		"p6",
		"p7",
		"p8",
		"p9",
	}

	mapInfo["map5"] = []string{
		"p1",
		"p2",
		"p3",
		"p4",
		"p5",
		"p6",
		"p7",
		"p8",
		"p9",
	}

	mapLen := len(mapInfo)
	fmt.Println(mapLen)
	fmt.Println(mapLen)
	fmt.Println(mapLen)
	pos := map[string]string{
		"map1_4": "map2_3",
		"map1_6": "map2_5",
		"map1_9": "map2_8",
		"map1_3": "map3_4",

		"map2_2": "map3_2",
		"map2_3": "map2_8",
		"map2_6": "map3_5",

		"map3_1": "map1_2",
		"map3_4": "map4_3",
		"map3_7": "map4_6",

		"map4_5": "map3_10",
		"map4_9": "map2_7",
		"map4_8": "map5_7",
	}
	NewMapArrive(mapInfo, pos)

	// arrrayMapArrive := make([][]bool, 1)
	// arrrayMapArrive[0][0] = true

	// 5个地图,每个地图有10个点
	//   期中有一些点会有连接 : 比如 map1p3 map2p2

	// fmt.Println(mapInfo)

	// strArr := []string{
	// 	"桃源村",
	// 	"桃源村",
	// }

	// arr := []Point{}
	// arr = append(arr, Point{
	// 	start: ""

	// 	"a01", "b01",
	// })

}
