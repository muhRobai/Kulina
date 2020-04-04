package main

import (
	"fmt"
)

// object declaration
type indonesianFood struct {
	Name string
	From string
}

//make func without input
func (i *indonesianFood) showFood() {
	fmt.Println("this indonesian food")
	var foodDesc string
	if i.Name != "" {
		foodDesc = fmt.Sprintf("%s", i.Name)	
	}

	if i.From != "" {
		foodDesc = fmt.Sprintf("%s From %s", foodDesc, i.From)
	}

	fmt.Println(foodDesc)
}

//func with input
func (i *indonesianFood) addIndonesianFood(name, from string) {
	i.Name = name
	i.From = from
}

func main() {
	//call the object
	c := indonesianFood{}
	//after call object you can call func in object
	c.showFood()
	c.addIndonesianFood("rendang", "padang")
	c.showFood()
	c.addIndonesianFood("soto", "")
	c.showFood()
}