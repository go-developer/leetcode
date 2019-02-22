package main

import "fmt"

func main() {
	n1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(n1))

	n2 := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(n2))

}

/**
 * 有序数组去重： 快慢指针法
 * @author go_developer@163.com <张德满>
 */
func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var (
		slowIndex int
		fastIndex int
	)

	slowIndex = 0
	for fastIndex = 1; fastIndex < len(nums); fastIndex++ {
		if nums[slowIndex] != nums[fastIndex] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]
		}
	}
	return slowIndex + 1
}
