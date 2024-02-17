package main

import "fmt"

func main() {
	type NoKTP string
	const ktpUserOne NoKTP = "1234567890"
	fmt.Println(ktpUserOne)

	//================================= with convertion
	var userNoKTP string = "0987654321"
	fmt.Println(NoKTP(userNoKTP))

	// or
	var anotherUserNoKTP string = "0987654321"
	var resultConvertKTP NoKTP = NoKTP(anotherUserNoKTP)
	fmt.Println(resultConvertKTP)
}
