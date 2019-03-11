package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	arrLen, firstStringLen, res := len(strs), len(strs[0]), []byte{}

	for i := 0; i < firstStringLen; i++ {
		for j := 1; j < arrLen; j++ {
			if len(strs[j]) <= i || strs[0][i] != strs[j][i] {
				return string(res)
			}
		}
		res = append(res, strs[0][i])
	}
	return string(res)
}

func main() {
	var k []string = []string{"flower","flow","flight"}

	fmt.Println(longestCommonPrefix(k))
}
