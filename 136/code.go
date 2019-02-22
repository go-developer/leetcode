package main

func main() {

}

/**
 * 异或运算的考察，同一个数异或偶数次，原数不变
 * @author go_developer@163.com <张德满>
 */
func singleNumber(nums []int) int {
	var (
		result int
	)
	result = 0
	for index := 0; index < len(nums); index++  {
		result = result ^ nums[index]
	}
	return result
}