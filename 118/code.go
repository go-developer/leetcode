package main

import "fmt"

func main() {
	fmt.Println(generate(5))

}

/**
 * 杨辉三角， 第几行就有几个数, 两端一定是1
 * @author go_developer@163.com <张德满>
 */
func generate(numRows int) [][]int {
	var (
		result [][]int
	)

	result = make([][]int, numRows)

	for row := 1; row <= numRows; row++ {
		tmp := make([]int, row)
		tmp[0] = 1
		tmp[row - 1] = 1
		if row <= 2 {
			result[row - 1] = tmp
			continue
		}

		for index := 1; index < row - 1; index++ {
			tmp[index] = result[row - 2][index - 1] + result[row - 2][index]
		}

		result[row - 1] = tmp
	}

	return result
}