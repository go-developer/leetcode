package main

import "strings"

func main()  {

}

/**
 * 计算给定的字符串列表中，所有字符串的最大公共前缀
 * 处理思路:
 * 1. 空list直接返回空字符串
 * 2. 只有一个元素的list直接返回唯一的元素
 * 3. 多有一个元素的字符串，先计算出长度最短的字符串，最长公共前缀的长度一定小于等于这个字符串，主要用于减少循环次数以及后续逻辑判断的复杂性
 * 4. 将列表中的字符串转为独立但自负的列表,此处为方便后续对字符的逐位比较
 * 5. 逐位比较所有字符串同一位的字符是否相同，此处引入一个临时map，若相同,则map长度一定是1,因为同key会覆盖,若长度不为1，终止比较，已比较过的字符即为最长公共字串
 * @author go_developer@163.com
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//查找最短的字符串的长度，最长公共前缀，长度一定小于等于该字符串
	minStrLength := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minStrLength > len(strs[i]) {
			minStrLength = len(strs[i])
		}
	}
	if minStrLength == 0 {
		//说明字符串列表中有空字符串
		return ""
	}

	strTable := make([][]string,0)
	for _, str := range strs {
		strArr := strings.Split(str, "")
		strTable = append(strTable, strArr[0:minStrLength])
	}

	commonPrefixList := make([]string, 0)
	for start := 0; start < minStrLength; start++ {
		tmpMap := make(map[string]bool)
		for i := 0; i < len(strTable); i++ {
			tmpMap[strTable[i][start]] = true
		}
		if len(tmpMap) == 1 {
			//所有字符串该位字符相同
			commonPrefixList = append(commonPrefixList, strTable[0][start])
		} else {
			break
		}
	}
	return strings.Join(commonPrefixList, "")
}
