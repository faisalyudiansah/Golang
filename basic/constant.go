package main

import "fmt"

const A = 1              // diluar scope, testing
var B string = "testing" // diluar scope, testing
var C int32 = 2          // diluar scope, testing

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	const firstname = "Cristiano"
	const lastname = "Ronaldo"
	const address = "Portugal" // const tanpa di gunakan, tidak akan error, tidak seperti variabel

	// error
	// firstname = "Lionel"
	// lastname = "Messi"

	fmt.Println(firstname)
	fmt.Println(lastname)

	fullName := firstname + " " + lastname // concatenation of strings
	fmt.Println("Full Name: ", fullName)

	//================================= Declare multiple constant
	const (
		playername = "Neymar"
		age        = 28
	)
	fmt.Println(playername)
	fmt.Println(age)

	const a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const ageCity, city = 55, "Madrid" // without declare data type
	country, ageCountry := "Spain", 78 // simplify

	fmt.Println(ageCity)
	fmt.Println(city)
	fmt.Println(country)
	fmt.Println(ageCountry)

}
