package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
	1 9 8
    x 2 9
	------------
		[2]
		[2 7]
		[2 3 2]
		[2 4 0 1]
		[2 4 8 2]
		[2 4 7 3]
		[2 4 7 5]
 */
func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}

	n1Len, n2Len := len(num1), len(num2)
	res := make([]int, 0, n1Len+n2Len)
	var n1, n2 = 0, 0

	for i := n1Len - 1; i >= 0; i-- {
		for j := n2Len - 1; j >= 0; j-- {
			n1, _ = strconv.Atoi(string(num1[i]))
			n2, _ = strconv.Atoi(string(num2[j]))

			index := (n1Len + n2Len - 2) - i - j
			res = deepFill(index, n1*n2, res)
		}
	}
	var buffer bytes.Buffer
	for i := len(res) - 1; i >= 0; i-- {
		buffer.WriteString(strconv.Itoa(res[i]))
	}

	return buffer.String()
}

/*
	res = [1, 9, 9, 0, 1]
	deepFill(1, 13, res) :
			[1,[ 9 + 13], 9, 0, 1]
			[1, 2, [2 + 9], 0, 1]
			[1, 2, 1, [1 + 0], 1]
			[1, 2, 1, 1, 1]
 */
func deepFill(index int, val int, res []int) []int {
	if index >= len(res) { //增加空间
		res = append(res, 0)
	}
	tmp := val + int(res[index])
	deciles := int(tmp / 10)

	res[index] = tmp % 10

	if deciles > 0 {
		res = deepFill(index+1, deciles, res)
	}

	return res
}

func main() {
	fmt.Println(multiply("123421412","1231321312321"))
}
