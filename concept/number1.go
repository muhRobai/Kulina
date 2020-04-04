package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

func convertArr(payload string) []int{
	//conver string array to int array
	var result []int
	value := strings.Split(payload, " ")
	for i := 0; i < len(value); i++ {
		val, err := strconv.Atoi(string(value[i]))
		if err != nil {
			fmt.Println(err)
			return result
		}
		result = append(result, val)
	}
	return result
}

func sortArr(arr []int) []int {
	var payload []int
	for i := 0; i < len(arr); i++ {
		for j := len(arr)-1; j >= i + 1; j-- {
			if int(arr[j]) < int(arr[j-1]){
				arr[j],arr[j-1] = arr[j-1],arr[j] // swap array
			}
		}
	}

	uniqValue := false
	for l := 0; l < len(arr); l++ {
		if l == 0 {
			payload = append(payload, int(arr[l]))
			continue
		}
		uniqValue = false
		for k := 0; k < len(payload); k++ {
			if int(arr[l]) == int(payload[k]) {
				uniqValue = true
			}
		}
		if uniqValue == false {
			payload = append(payload, int(arr[l]))
		}
	}
	return payload
}
 
func MatchingColor(number int, payload string) (int, error){
	arr := convertArr(payload)
	if len(arr) != number {
		fmt.Println("invalid-input")
		return 0, errors.New("invalid-input")
	}

	sortArr := sortArr(arr)
	var item []int 
	// var colorDone []string
	var countColor int
	 var color int
	for i := 0; i < len(sortArr); i++ {
		countColor = 0
		color = sortArr[i]
		for l := 0; l < len(arr) ; l++  {
			if color == int(arr[l]) {
				countColor++
			}
		}

		item = append(item, countColor)
	}

	var result int
	for i := 0 ; i < len(item); i++ {
		result += int(item[i])/2
	}

	return result, nil
 }

func main() {
	fmt.Print("> ")
	resp, err := MatchingColor(9, "10 20 20 10 10 30 50 10 20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}