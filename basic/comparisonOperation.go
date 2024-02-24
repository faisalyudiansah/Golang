package main

import (
	"fmt"
	"reflect"
)

func main() {
	const a = 2
	const b = 3
	var c bool = a == b
	fmt.Println(c)

	//====================================== comparison bool
	var isGood bool = a < b
	fmt.Println(isGood, "<<<<<<<<<<")

	//====================================== reflect / type of comparison
	userOne := "Ronaldo"
	var check bool = reflect.TypeOf(userOne).Kind() == reflect.String
	fmt.Println(check)
	if reflect.TypeOf(userOne).Kind() == reflect.String {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
