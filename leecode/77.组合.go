package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	rList := make([]int, 0, 100)
	res := deepCombine(n, k, 1, rList, [][]int{})

	return res
}
func deepCombine(n, k, start int, list []int, res [][]int) [][]int {
	if len(list) == k {
		v := make([]int, 0, len(list))
		v = append(v, list...)
		res = append(res, v)
		return res
	}
	for i := start; i <= n - (k - len(list)) + 1; i++ {
		list = append(list, i)
		res = deepCombine(n, k, i+1, list, res)
		list = list[:len(list)-1]
	}

	return res
}