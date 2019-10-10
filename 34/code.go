package main

import "fmt"

func main() {
	testCase := []int{5,7}
	fmt.Println(searchRange(testCase, 7))
}

/**
 * @author go_developer@163.com
 */
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	begin := 0
	end := len(nums) - 1
	startIndex := -1
	endIndex := -1
	for  {
		middleIndex := (begin + end) / 2
		if nums[middleIndex] > target {
			end = middleIndex - 1
		} else if nums[middleIndex] < target {
			begin = middleIndex + 1
		} else {
			//计算起始索引
			for i := middleIndex; i >= begin; i-- {
				if nums[i] == target {
					startIndex = i
				} else {
					break
				}
			}
			//计算终止索引
			for i := middleIndex; i <= end; i++ {
				if nums[i] == target {
					endIndex = i
				} else {
					break
				}
			}
			break
		}
		if begin > end {
			break
		}
	}
	return []int{startIndex, endIndex}
}
