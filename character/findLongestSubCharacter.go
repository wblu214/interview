package main

import (
	"fmt"
)

/*给定一个字符串 s，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1：
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2：
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3：
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。*/

func lengthOfLongestSubstring(s string) int {
	// 使用map来记录字符最后出现的位置
	lastOccurred := make(map[byte]int)

	start := 0     // 窗口起始位置
	maxLength := 0 // 最大长度

	for i, ch := range []byte(s) {
		// 如果字符已经出现过，并且位置在当前窗口内
		if last, ok := lastOccurred[ch]; ok && last >= start {
			// 移动窗口起始位置到重复字符的下一个位置
			start = last + 1
		}
		// 更新字符最后出现的位置
		lastOccurred[ch] = i
		// 计算当前窗口长度，更新最大值
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}

func main() {
	//ww我
	source := "abcabcbb"
	fmt.Printf("原字符串为：%s\n", source)
	result := lengthOfLongestSubstring(source)
	fmt.Printf("最大不重复子串长度为：%d\n", result)
}
