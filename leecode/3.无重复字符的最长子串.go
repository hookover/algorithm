package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(str string) int {
	stringMap, longest, start := map[int32]int{}, 0, 0
	utf8Chars := []rune(str)
	for index, char := range utf8Chars {
		if charIdx, ok := stringMap[char]; ok && start <= charIdx{
			start = charIdx + 1
		} else {
			longest = int(math.Max(float64(longest), float64(index - start + 1)))
		}
		stringMap[char] = index
	}

	return longest
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
