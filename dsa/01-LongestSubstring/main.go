package main

import "fmt"

func LongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	lastIndex := make([]int, 128)

	start := 0
	for i := 0; i < n; i++ {
		currentChar := s[i]

		if lastIndex[currentChar] > start {

			start = lastIndex[currentChar]

		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastIndex[currentChar] = i + 1
	}
	return maxLength
}

func main() {
	str := "amaanmirza"
	val1 := LongestSubstring(str)
	fmt.Println(val1)
}
