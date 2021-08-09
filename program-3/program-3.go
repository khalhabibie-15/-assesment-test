package main

import (
	"fmt"
)

func mergerString(string1, string2 string) string {
	lengthString1 := len(string1)
	lengthString2 := len(string2)
	var result string

	if lengthString1 < lengthString2 {
		for i := 0; i < lengthString1; i++ {
			result += string(string1[i])
			result += string(string2[i])
		}
		result += string(string2[lengthString1:lengthString2])
	} else if lengthString1 > lengthString2 {
		for i := 0; i < lengthString2; i++ {
			result += string(string1[i])
			result += string(string2[i])
		}
		result += string1[lengthString2:lengthString1]
	} else {
		for i := 0; i < lengthString2; i++ {
			result += string(string1[i])
			result += string(string2[i])
		}
	}
	return result

}

func main() {

	string1 := "saya"
	string2 := "kamu"
	fmt.Printf("String 1 : %s\nString 2 : %s\nResult  : %s \n", string1, string2, mergerString(string1, string2))

	string1 = "kosong"
	string2 = "ada"
	fmt.Printf("\nString 1 : %s\nString 2 : %s\nResult  : %s\n", string1, string2, mergerString(string1, string2))

	string1 = "ada"
	string2 = "kosong"
	fmt.Printf("\nString 1 : %s\nString 2 : %s\nResult  : %s", string1, string2, mergerString(string1, string2))

}
