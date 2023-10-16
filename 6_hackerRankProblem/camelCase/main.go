package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	count := 1
	for _, char := range input {
		str := string(char)
		//verify str is already uppercase
		if strings.ToUpper(str) == str {
			count++
		}
	}
	fmt.Println(count)
}
