package main

import (
	"fmt"
	"strings"
	"unicode"
)

func stackLetter(a, b string) string {
	a += "z"
	b += "z"
	var resultSlice []string

	length := len(a) + len(b) - 2
	for i := 0; i < length; i++ {
		if a < b {
			resultSlice = append(resultSlice, string(a[0]))
			a = a[1:]
		} else {
			resultSlice = append(resultSlice, string(b[0]))
			b = b[1:]
		}
	}

	return strings.Join(resultSlice, "")

}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	resultSlice := []string{}
	var string1 string
	var string2 string

	i := 0
	for i < n {
		fmt.Scanf("%v \n %v", &string1, &string2)
		if !IsUpper(string1) || !IsUpper(string2) {
			fmt.Println("WARNING : Please input upper case")
			return
		}
		resultString := stackLetter(string1, string2)
		resultSlice = append(resultSlice, resultString)
		i++
	}

	for _, string := range resultSlice {
		fmt.Println(string)
	}

}
