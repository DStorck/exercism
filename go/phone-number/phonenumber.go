package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(str string) (string, error) {
	reg, _ := regexp.Compile("[^0-9']+")
	num := reg.ReplaceAllString(str, "")
	if len(num) > 11 || len(num) < 10 {
		return "", errors.New("Invalid length")
	}
	if len(num) == 11 && num[0:1] == "1" {
		num = num[1:]
	}
	if len(num) != 10 {
		return "", errors.New("not ten digits ")
	}
	if num[0:1] == "1" || num[3:4] == "1" || len(num) == 11 || num[0:1] == "0" || num[3:4] == "0" {
		return "", errors.New("BADNESS 1000")
	}
	return num, nil

}

func Format(str string) (string, error) {
	str, err := Number(str)
	if err == nil {
		s := fmt.Sprintf("(%v) %v-%v", str[0:3], str[3:6], str[6:])
		return s, nil
	}
	return "", err
}

func AreaCode(str string) (string, error) {
	code, err := Number(str)
	if err == nil {
		return code[0:3], nil
	}
	return "", err
}
