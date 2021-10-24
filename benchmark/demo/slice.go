package demo

func appendOne(num int) []int {
	var res []int
	for i := 0; i < num; i++ {
		res = append(res, i)
	}
	return res
}

func appendMany(num int) []int {
	res := make([]int, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, i)
	}
	return res
}
