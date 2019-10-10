package main

import "fmt"

/**
 * 盛水最多的容器
 * @author go_developer@163.com
 */
func main()  {
	var testCase = []int{2,3,4,5,18,17,6}
	fmt.Println(maxArea(testCase))
}

/**
 * 双指针法计算最大容量
 * @author go_developer@163.com
 */
func maxArea(height []int) int {
	left := 0	//左起点
	right := len(height) - 1 //右起点
	maxVal := 0 //最大容量
	for left != right {
		isLeftSmall := true //是否左边较小
		currentHeight := height[left]
		if currentHeight > height[right] {
			isLeftSmall = false
			currentHeight = height[right]
		}
		tmpVal := (right - left) * currentHeight
		if tmpVal > maxVal {
			maxVal = tmpVal
		}

		if isLeftSmall {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return maxVal
}
