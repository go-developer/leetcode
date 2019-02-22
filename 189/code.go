package main

import (
	"fmt"
)

func main() {
	rotate([]int{1, 2, 3,4,5,6,7}, 3)
}

/**
 * 本方案： 不需要任何中间数组，直接在原数组之上进行操作, (应该还有更优解)
 * @author go_developer@163.com <张德满>
 */
func rotate(nums []int, k int) {
	var (
		numLen int
		//移动位数
		moveBit int
	)

	numLen = len(nums)
	moveBit = k % numLen
	if numLen <= 1 {
		return
	}

	for c := 1; c <= moveBit; c++  {
		removeVal := nums[0]
		for index := 0; index < numLen; index++ {
			//移动一位后的索引
			lastIndex := (index + 1) % numLen
			//目标索引的原值
			lastIndexVal := nums[lastIndex]
			//每次向后移动一位
			nums[lastIndex] = removeVal
			removeVal = lastIndexVal
		}
	}

	fmt.Println(nums)
}
