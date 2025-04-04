// Firsts steps with variables and typing
package main

import "fmt"

func main() {

	// Strings Section
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"
	nameFour := "Yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Ints Section
	var ageOne int = 20
	var ageTwo = 17
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)
}
