package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
}

/**
 * 数字搜索，因为有序，直接二分查找
 * @author go_developer@163.com <张德满>
 */
func searchInsert(nums []int, target int) int {
	var (
		low int
		high int
		middle int
	)
	low = 0
	high = len(nums)

	for low < high {
		middle = low + (high - low) / 2
		if nums[middle] > target {
			high = middle
		} else if nums[middle] < target {
			low = middle + 1
		} else {
			return middle
		}
	}

	return low
}
