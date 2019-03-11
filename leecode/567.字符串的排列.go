package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	s1Arr, s2Arr := [26]int{}, [26]int{}
	s1Len, s2Len := len(s1), len(s2)

	if s1Len > s2Len {
		return false
	}

	for i := 0; i < s1Len; i++ {
		s1Arr[s1[i]-'a'] ++
		s2Arr[s2[i]-'a'] ++
	}

	if s1Arr == s2Arr {
		return true
	}

	for i := s1Len; i < s2Len; i++ {
		s2Arr[s2[i]-'a'] ++
		s2Arr[s2[i-s1Len]-'a'] --

		if s1Arr == s2Arr {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab","eidbaooo"))
}