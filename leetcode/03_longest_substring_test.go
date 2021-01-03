package leetcode

import "testing"

func TestLongestSubstring(t *testing.T) {
	longest := lengthOfLongestSubstring("abcbacade")
	if longest != 3 {
		t.Errorf("failed , longest %v", longest)
	}
}
