package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9 ,9}))
}


/**
 * 留意最高一位进1情况
 * @author go_developer@163.com <张德满>
 */
func plusOne(digits []int) []int {
	var (
		digitLen int
		carry int
	)

	digitLen = len(digits)

	carry = 0
	if (digits[digitLen - 1] + 1) >= 10 {
		carry = 1
	}

	digits[digitLen - 1] = (digits[digitLen - 1] + 1) % 10

	for index := digitLen - 2; index >= 0; index-- {
		val := digits[index] + carry
		digits[index] = val % 10
		if val >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}

	//说明书组要扩容了 如 999 + 1 = 1000 原数组容量只有三
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
