package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func isValidChar(s string) bool {
	isokay := regexp.MustCompile(`[0-9\s]+`).MatchString
	return isokay(s)
}

func createNumArray(str string) []int {
	numArr := make([]int, 0)
	chars := strings.Split(str, "")
	for _, char := range chars {
		num, err := strconv.Atoi(char)
		is_okay := isValidChar(char)
		if err == nil {
			numArr = append(numArr, num)
		}
		if is_okay == false {
			numArr = []int{1}
			break
		}
	}
	return numArr
}

func doubleSecond(arr []int) []int {
	count := 1
	for i := len(arr) - 2; i > -1; i-- {
		if count%2 == 0 {
			count += 1
			continue
		}
		count += 1
		temp := arr[i] * 2
		if temp > 9 {
			arr[i] = temp - 9
		} else {
			arr[i] = temp
		}
	}
	return arr
}

func checkSum(arr []int) bool {
	var sum int
	for _, num := range arr {
		sum += num
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func Valid(s string) bool {
	a := createNumArray(s)
	if len(a) == 1 {
		return false
	}
	b := doubleSecond(a)
	c := checkSum(b)
	return c
}
