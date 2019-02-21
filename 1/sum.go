package main

import "fmt"

func main()  {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

/**
 * 实现方式一共经过两次list遍历，可以考虑是否能缩减至一次
 * @author go_developer@163.com <张德满>
 */
func twoSum(nums []int, target int) []int {
	var (
		listLen int
		result []int
		numsMap map[int]int
		expectAnthorValue int
	)

	listLen = len(nums)
	result = make([]int, 0)

	if listLen < 1 {
		return result
	}

	if 1 == listLen {
		if target == nums[1] {
			result = append(result, 0)
		}
		return result
	}

	//将numbers转换为map,方便后续处理
	numsMap = make(map[int]int)
	for index := 0; index < listLen; index++ {
		numsMap[nums[index]] = index
	}

	//寻找数值对
	for index := 0; index < listLen - 1; index++ {
		expectAnthorValue = target - nums[index]
		if _, ok := numsMap[expectAnthorValue]; ok {
			if index != numsMap[expectAnthorValue] {
				//找到了数值组(主义同一个索引数据不能重复使用)
				result = append(result, index)
				result = append(result, numsMap[expectAnthorValue])
				return result
			}
		}
	}
	return result
}