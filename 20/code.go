// Package _0...
//
// Description : _0...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-09 10:26 下午
package main

import (
	"fmt"
)

func main() {
	testCase := []string{
		"()",
		"()[]{}",
		"([)]",
		"(]",
	}

	for _, itemCase := range testCase {
		fmt.Printf("%v \n", isValid(itemCase))
	}
}

// isValid 判断是否为有效的字符串
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:57 下午 2021/2/9
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	// 括号标示的数据结构
	type charFlag struct {
		matchChar string
		isRight   bool
	}

	charTable := map[string]charFlag{
		"(": {isRight: false},
		")": {isRight: true, matchChar: "("},
		"[": {isRight: false},
		"]": {isRight: true, matchChar: "["},
		"{": {isRight: false},
		"}": {isRight: true, matchChar: "{"},
	}

	waitMatchCharList := make([]string, 0)
	for _, c := range s {
		var (
			exist      bool
			charConfig charFlag
		)
		if charConfig, exist = charTable[string(c)]; !exist {
			continue
		}
		if !charConfig.isRight {
			// 左括号入栈
			waitMatchCharList = append(waitMatchCharList, string(c))
			continue
		}
		// 右括号匹配
		if len(waitMatchCharList) == 0 {
			// 栈已空，匹配失败
			return false
		}
		// 判断栈顶和当前括号是否匹配
		if waitMatchCharList[len(waitMatchCharList)-1] != charConfig.matchChar {
			// 括号不匹配
			return false
		}
		waitMatchCharList = waitMatchCharList[:len(waitMatchCharList)-1]
	}
	if len(waitMatchCharList) == 0 {
		// 完全匹配
		return true
	}

	return false
}
