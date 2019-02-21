package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("abb"))
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("abcd"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

/**
 * @author go_developer@163.com <张德满>
 */
func lengthOfLongestSubstring(s string) int {
	var (
		strLen int
		maxLength int
		charArr []string
		startIndex int
		index int
	)

	maxLength = 0
	startIndex = 0

	strLen = len(s)
	if strLen <= 1 {
		//空串或长度为一的字符串
		return strLen
	}

	charArr = strings.Split(s, "")

	for {
		if startIndex >= strLen {
			break
		}

		tmpLen := 0

		// key 字符 val 字符所处的索引
		charTable := make(map[string]int)
		//重复的字符
		repeatChar := ""

		for index = startIndex; index < strLen; index++ {
			if _, ok := charTable[charArr[index]]; ok {
				//字符已重复
				repeatChar = charArr[index]
				break
			}
			charTable[charArr[index]] = index
			tmpLen = tmpLen + 1
		}

		startIndex = charTable[repeatChar] + 1

		if tmpLen > maxLength {
			//当前子串长度大于上一个字串长度
			maxLength = tmpLen
		}

		if (maxLength + startIndex) >= strLen {
			//说明后面的字串，即使全不重复，最大长度也和当前一样，无需继续迭代
			break
		}
	}

	return maxLength
}