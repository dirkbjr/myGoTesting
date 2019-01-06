// Ninja Level 12 - Hands-On Exercise from Udemy course on Learning Go Programming,
// by Todd McLeod. Creating documentation in the code and having it referenced on godoc.org

//Package main - ...
//I felt like doing something different than the dog years example in the course syllabus :-)
package main

import (
	"fmt"

	"github.com/myGoTesting/01-udemy-handson-12-1/petty"
)

func main() {
	var carModel string
	var modelYear string //I might change this to int in the future

	//The user inputs one of three car model types, Camaro, Firebird, Mustang
	fmt.Print("Enter Car Model Name (Camaro | Firebird | Mustang): ")
	_, err := fmt.Scan(&carModel)
	if err != nil {
		panic(err)
	}

	//The user inputs either 1996 or 1997, as the year of model sales
	fmt.Print("Enter a specific Model Year between 1996-1997: ")
	_, err = fmt.Scan(&modelYear)
	if err != nil {
		panic(err)
	}

	//Now that you have the desired car model and model year, call the petty function to print out
	//the desired statistics
	petty.OutputSales(carModel, modelYear)

}
