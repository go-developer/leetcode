package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(reverse(-123))
	fmt.Println(reverse(123))
	fmt.Println(reverse(120))

}

/**
 * 注意数据溢出检查
 * @author go_developer@163.com <张德满>
 */
func reverse(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		//按题目规定，输入数据已溢出
		return 0
	}
	if x > -10 && x < 10 {
		return x
	}

	var (
		numList []int
		mod int
		result int
		topNum int
		tmp int
	)

	if x >= 0 {
		topNum = math.MaxInt32
	} else {
		topNum = math.MinInt32
	}
	numList = make([]int, 0)
	tmp = x
	for  {
		mod = tmp % 10
		tmp = tmp / 10
		numList = append(numList, mod)
		if tmp == 0 {
			break
		}

	}
	numLen := len(numList)

	result = 0

	for index := 0; index < numLen; index++ {
		lastNum := numList[index] * int(math.Pow10(numLen - 1 - index))
		if x >= 0  {
			if (topNum - lastNum) < result {
				//溢出
				return 0
			} else {
				result = result + lastNum
			}
		} else {
			if (topNum - lastNum) > result {
				//溢出
				return 0
			} else {
				result = result + lastNum
			}
		}

	}

	return result
}