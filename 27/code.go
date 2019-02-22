package main

import "fmt"

func main() {
	n1 := []int{1, 1, 2}
	fmt.Println(removeElement(n1, 1))

	n2 := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeElement(n2, 1))

}

/**
 * 有序数组移除元素： 快慢指针法
 * @author go_developer@163.com <张德满>
 */
func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	var (
		slowIndex int
		fastIndex int
	)

	slowIndex = 0
	for fastIndex = 0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}