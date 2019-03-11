package main

import "fmt"

/*
	abc	1234
	4321 cba 翻转

	i  j 交换
	1324
	1234
		 i j 交换
		 abc
	1234 abc
 */
func reverseWords(s string) string {
	strArr := []rune(s)

	strArr = deepSection(0, strArr)
	strHalfLen := len(strArr) / 2

	for i, j := 0, len(strArr)-1; i < strHalfLen; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}

	//fmt.Println(string(strArr))

	strLen := len(strArr)
	for i, j := 0, 0; j < strLen; j++ {
		if string(strArr[j]) == " " || j == strLen-1 {
			if j == strLen-1 {
				swap(strArr, i, j)
			} else {
				swap(strArr, i, j-1)
			}
			i = j + 1
			j = i
		}
	}

	return string(strArr)
}
func swap(arr []rune, i, j int) {
	if i == j {
		return
	}

	half := (j-i + 1)/2 + i

	for ; i < half; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func deepSection(index int, slice []rune) []rune {
	strLen := len(slice)
	if index >= len(slice) {
		return slice
	}

	if string(slice[index]) == " " && ((index == 0 || index == strLen-1) || (index+1 < strLen && string(slice[index+1]) == " ")) {
		slice = append(slice[:index], slice[index+1:]...)
		index = index - 1 //删除后index会缩减
	}

	if index+1 < len(slice) {
		slice = deepSection(index+1, slice)
	}

	return slice
}

func main() {
	fmt.Println(reverseWords("a good   example "))
}
