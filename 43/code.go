package main

import (
	"strconv"
	"strings"
)

func main()  {
	multiply("99", "999")
}

/**
 * 字符串相乘(大数相乘)
 * @author go_developer@163.com
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Arr := strings.Split(num1, "")
	num2Arr := strings.Split(num2, "")
	//将number2转化为int的list
	num2IntList := make([]int, 0)
	for i := 0; i < len(num2Arr); i++ {
		val, _ := strconv.Atoi(num2Arr[i])
		num2IntList = append(num2IntList, val)
	}

	var result  [111][]int
	//计算num1与num2的乘积
	for i := 0; i < len(num1Arr); i++ {
		//计算后，末尾需要补0的数量
		needZeroCnt := len(num1Arr) - 1 - i
		intVal, _ := strconv.Atoi(num1Arr[i])
		tmpResult := make([]int, 0)
		lastExtraAdd := 0
		//此步计算结果低位在前，高位在后
		for start := len(num2Arr) - 1; start >= 0; start-- {
			val := intVal * num2IntList[start] + lastExtraAdd
			leave := val % 10	//保留的数字
			lastExtraAdd = (val - leave) / 10 //进位的值
			tmpResult = append(tmpResult, leave)
		}
		if lastExtraAdd > 0 {
			tmpResult = append(tmpResult, lastExtraAdd)
		}
		//倒序，转换成正常的高位在前，低位在后顺序
		formatResult := make([]int, 0)
		for start := len(tmpResult) - 1; start >= 0; start-- {
			formatResult = append(formatResult, tmpResult[start])
		}
		result[i] = append(result[i], formatResult...)
		//补齐末尾应有的0
		for cnt := 0; cnt < needZeroCnt; cnt++ {
			result[i] = append(result[i], 0)
		}
		//对结果高位以0补位对齐
		highZeroCnt := len(result[0]) - len(result[i])
		zeroList := make([]int, 0)
		for i := 0; i < highZeroCnt; i++ {
			zeroList = append(zeroList, 0)
		}
		result[i] = append(zeroList, result[i]...)
	}

	//对结果求和
	sumResult := make([]int, 0)
	lastExtraAdd := 0
	for start := len(result[0]) - 1; start >= 0; start-- {
		tmpSum := lastExtraAdd
		for i := 0; i < len(num1Arr); i++ {
			tmpSum = tmpSum + result[i][start]
		}
		leaveSum := tmpSum % 10
		lastExtraAdd = (tmpSum - leaveSum) / 10
		sumResult = append(sumResult, leaveSum)
	}
	if lastExtraAdd > 0 {
		sumResult = append(sumResult, lastExtraAdd)
	}

	outputStr := ""
	for start := len(sumResult) - 1; start >= 0; start-- {
		outputStr = outputStr + strconv.Itoa(sumResult[start])
	}
	return outputStr
}
