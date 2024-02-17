package main

import "fmt"

func main() {
	//====================================== string
	fmt.Println(`string 1`) // use ` `
	fmt.Println("string 2") // use " " and cannot use ' '
	fmt.Println("")         // empty

	//====================================== function string
	fmt.Println(len("string"))         // len , like length in javascript
	fmt.Println("Faisal Yudiansah"[0]) // take the character at index 4, it will return "byte" / uint8. Convert to string if want "F"

	//====================================== number int and float
	fmt.Println(`int one = `, 1) // use ` `
	fmt.Println("int two = ", 2) // use " " and cannot use ' '
	fmt.Println("float = ", 2.5) // float

	//====================================== bool
	fmt.Println("bool true", true)
	fmt.Println("bool false", false)
}
