package main

import "fmt"

func simplifyPath(path string) string {
	slice := make([]string, 0)
	pathArr := []rune(path)
	strLen := len(pathArr)
	m, n := 0, 0
	var resString string

	for i, j := 0, 0; j < strLen; j++ {
		if i < j && (string(pathArr[j]) == "/" || j == strLen-1) {
			m, n = i, j

			if string(pathArr[i]) == "/" {
				m ++
			}
			if j == strLen-1 && string(pathArr[j]) != "/" {
				n++
			}

			i = j

			if m < n {
				switch string(path[m:n]) {
				case ".":
				case "..":
					if len(slice) >= 1 {
						slice = slice[:len(slice)-1]
					}
				default:
					slice = append(slice, string(path[m:n]))
				}
			}
		}
	}

	for _, str := range slice {
		resString = resString + "/" + str
	}
	if len(resString) == 0 {
		resString = "/"
	}

	return resString
}

func main() {
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
}
