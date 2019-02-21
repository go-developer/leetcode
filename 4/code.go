package main

import "fmt"

func main() {
	data := []int{1, 3}
	data1:= []int{2}
	fmt.Println(findMedianSortedArrays(data, data1))

	data = []int{1, 3}
	data1= []int{2, 4}
	fmt.Println(findMedianSortedArrays(data, data1))
}

/**
 * 合并数组后，排序取中位数(时间复杂度比较高，等同于所选择的排序算法的时间复杂度)
 * @author go_developer@163.com <张德满>
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		nums []int
	)
	nums = append(nums1, nums2...)

	if len(nums) == 1 {
		return float64(nums[0])
	}

	if len(nums) == 2 {
		return (float64(nums[0]) + float64(nums[1])) / 2
	}

	//排序，升序降序都可以(算法 ： 插入排序, 升序)
	for index := 1; index < len(nums); index++ {
		for tmpIndex := 0; tmpIndex < index; tmpIndex++ {
			tmpVal := nums[index]
			if tmpVal < nums[tmpIndex] {
				//插入相关位置(tmpIndex, index之间的数据后移)
				for reIndex := index; reIndex > tmpIndex; reIndex-- {
					nums[reIndex] = nums[reIndex - 1]
				}
				nums[tmpIndex] = tmpVal
			}
		}
	}

	if len(nums) % 2 == 0 {
		//偶数个元素
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2 -1])) / 2
	} else {
		//奇数个元素
		return float64(nums[(len(nums) - 1)/2])
	}
}