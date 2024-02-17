package main

import "fmt"

func main() {
	//====================================== Convert data type
	var nilai32 int32 = 32768
	fmt.Println(nilai32)
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	const name = "Faisal Yudiansah"
	fmt.Println("Byte : ", name[2]) // byte or uint8
	fmt.Println(string(name[2]))    // yang awalnya "byte" di convert ke string
	fmt.Println(string(nilai64))    // wtf is this
}
