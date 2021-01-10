// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package leetcode

import "fmt"

/*
Runtime: 248 ms, faster than 17.37% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 6.5 MB, less than 20.61% of Go online submissions for Longest Substring Without Repeating Characters.
*/
func lengthOfLongestSubstringx(s string) int {
	r := []rune(s) //abcdabc
	rLen := len(r)
	count := 0
	maxCount := count

	for i := 0; i < rLen; i++ {

		maps := make(map[rune]bool)
		maps[r[i]] = true

		count++

		for j := i + 1; j < rLen; j++ {

			if r[i] == r[j] {
				break
			}

			if _, ok := maps[r[j]]; ok {
				break
			}

			maps[r[j]] = true

			count++
		}

		if count > maxCount {
			maxCount = count
		}

		count = 0

	}

	return maxCount
}

//TBC use sliding window
func lengthOfLongestSubstring(s string) int {
	rn := []rune(s)
	maps := make(map[rune]int)
	counter := 0
	starter := 0
	maxLength := 0

	//abcabcd
	for i, r := range rn {

		if val, ok := maps[r]; !ok {

			maps[r] = i
			counter++

			fmt.Println("rune ", r, " idx ", i, " c ", counter)

		} else {

			fmt.Println("starter ", starter, " val ", val, " i ", i)
			if starter <= val && val <= i {
				starter = i
				counter = 1
			} else {
				counter++
			}

			maps[r] = i
			fmt.Println("else rune ", r, " idx ", i, " c ", counter)

		}

		if counter > maxLength {
			maxLength = counter
		}
	}

	return maxLength
}
