package main

import (
	"fmt"
	"strings"
)

func generateNumber(payload string) {
	str := strings.ReplaceAll(payload, ".", "")
	arr := strings.Split(str, "")
	lengthArr := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		for l := 0; l < lengthArr; l++ {
			fmt.Print("0")
		}
		lengthArr--
		fmt.Println("")
	}
}

func main() {
	generateNumber("1.123.121")
}