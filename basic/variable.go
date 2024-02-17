package main

import "fmt"

func main() {
	// var name string // jika tanpa data, tipe data nya tulis
	// name = "Cristiano Ronaldo"
	// fmt.Println(name)

	// name = "Lionel Messi"
	// fmt.Println(name)
	// fmt.Println(len(name))

	// var club = "Real Madrid"
	// fmt.Println(club)
	// println(club) // can use this, but what the diffrent with fmt.Println() ?

	// money := "rupiah" // inferred / simplify (tanpa tulis var)
	// fmt.Println(money)
	// println(money)

	//================================= Declare multiple variable
	var (
		playername = "Neymar"
		age        = 28
	)
	fmt.Println(playername)
	fmt.Println(age)

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var ageCity, city = 55, "Madrid"
	country, ageCountry := "Spain", 78

	fmt.Println(ageCity)
	fmt.Println(city)
	fmt.Println(country)
	fmt.Println(ageCountry)
}
