package dira07

import (
	"fmt"
	"image/color"
	"math"
	"testing"
)

// [匿名组合](https://www.jb51.net/article/146333.htm)
func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

// 以上代码可以理解为自动生成: 内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}
func TestYxx(t *testing.T) {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"

	// p.Distance(q)  // 编译错误
	p.Distance(q.Point) //
}
