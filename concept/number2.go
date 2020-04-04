package main

import (
	"fmt"
	"strings"
	"errors"
)

func countClimb(numb int, payload string) (int, error) {
	arr := strings.Split(payload, " ")

	if len(arr) != numb {
		return 0, errors.New("invalid-input")
	} 
	
	countClimb := 0
	result := 0
	changes := true
	for i := 0; i < len(arr); i++ {
		if string(arr[i]) == "D" {
			countClimb--
		} else {
			countClimb++	
		}

		if countClimb == 0 {
			changes = true
		}

		if countClimb < 0 && changes == true {
			result++
			changes = false
		}
	}

	return result, nil
	
}	

func main(){
	resp, err := countClimb(8, "U D D D U D U U")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}