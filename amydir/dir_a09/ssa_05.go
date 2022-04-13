package dira09

func M05() []int {
	s := []int{3, 2, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	i := 5
	c := append(s[:i], s[i+1:]...) // 最后面的“...”不能省略
	return c
}
