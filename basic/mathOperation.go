package main

import "fmt"

func main() {
	var a = 12
	var b = 2
	var pengurangan = a - b
	var modulus = a % b
	fmt.Println("pengurangan", pengurangan)
	fmt.Println("pertambahan", a+b)
	fmt.Println("modulus", modulus)

	//================================= augmented assignment += -= *= /= %=
	x := 0
	x += 2
	fmt.Println(x)

	// OR

	v := 0
	v = v + 2
	fmt.Println(v)

	//================================= += -= *= /= %=

	z := 0
	z -= 5
	fmt.Println(z)

	f := 12
	f %= 2
	fmt.Println(f)
}
