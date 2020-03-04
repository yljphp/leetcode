package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "{[]}()[]{}"

	fmt.Println(isValid(str))
}

func isValid(s string) bool {

	tmpSlice := make([]byte, 0)

	for _, v := range s {

		if v == '{' || v == '(' || v == '[' {

			tmpSlice = append(tmpSlice, byte(v))

		} else {
			if len(tmpSlice) <= 0 {
				return false
			}

			if v == '}' && tmpSlice[len(tmpSlice)-1:][0] != '{' {
				return false
			}
			if v == ']' && tmpSlice[len(tmpSlice)-1:][0] != '[' {
				return false
			}
			if v == ')' && tmpSlice[len(tmpSlice)-1:][0] != '(' {
				return false
			}

			tmpSlice = tmpSlice[:len(tmpSlice)-1]
		}
	}

	return len(tmpSlice) == 0
}

//最简单，但是执行效率太差了
func isValidBak(s string) bool {

	for {

		if !strings.Contains(s, "{}") && !strings.Contains(s, "[]") && !strings.Contains(s, "()") {
			break
		}

		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "()", "")
	}

	if strings.TrimSpace(s) == "" {
		return true
	}

	return false
}
