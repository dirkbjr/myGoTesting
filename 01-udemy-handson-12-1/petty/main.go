// Ninja Level 12 - Hands-On Exercise from Udemy course on Learning Go Programming,
// by Todd McLeod. Creating documentation in the code and having it referenced on godoc.org

//Package petty - Provides basic statistics on car sales for traditional U.S. muscle cars
//I felt like doing something different than the dog years example in the course syllabus :-)
package petty

import (
	"encoding/json"
	"fmt"
)

//For this little test program, the source of data is carsalesbase.com. Model year sales for
//1996 and 1997 were placed in a slice of bytes, which will be json.UnMarshaled into a Go
//data structure
var jsonBlob = []byte(`[
	{"CarName": "Camaro", "CarYear": "1996", "CarSales": "66866"},
	{"CarName": "Camaro", "CarYear": "1997", "CarSales": "55973"},
	{"CarName": "Firebird", "CarYear": "1996", "CarSales": "32622"},
	{"CarName": "Firebird", "CarYear": "1997", "CarSales": "32524"},
	{"CarName": "Mustang", "CarYear": "1996", "CarSales": "122674"},
	{"CarName": "Mustang", "CarYear": "1997", "CarSales": "116610"}
	]`)

type carModelSales struct {
	CarName  string
	CarYear  string
	CarSales string
}

var allCars []carModelSales

//OutputSales takes a Car Model (carModel) and Model Year (modelYear) as input,
//populates the vehicle sales database variable, finds the desired Car Model and Model Year,
// and then prints out the full name of the vehicle and the sales for the desired Model Year.
func OutputSales(carModel, modelYear string) {
	var fullName, carSalesOut string

	//populateDb function is called to take the slice of bytes and place it into an array of struct
	populateDb()

	//Print out what the user input was
	switch carModel {
	case "Camaro":
		fullName = "Chevrolet Camaro"
	case "Firebird":
		fullName = "Pontiac Firebird"
	case "Mustang":
		fullName = "Ford Mustang"
	}

	fmt.Println("")

	for _, v := range allCars {
		if v.CarName == carModel {
			if v.CarYear == modelYear {
				carSalesOut = v.CarSales
			}
		}
	}
	//Print out the Full Name of the vehicle, the model year chosen and the total sales for the
	//specified year
	fmt.Println("\t", fullName, "\tSales in Model Year ", modelYear, "\t-> ", carSalesOut)
	fmt.Println("")
}

//populateDb() function takes the global variable jsonBlob, which is a slice of bytes comprised of
//json formatted data for the car models and their sales for the specified years. The json.Unmarshal function
//is called to populate a Go variable allCars, which is an array of struct carModelSales
func populateDb() {

	err := json.Unmarshal(jsonBlob, &allCars)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Printf("%+v\n", allCars)

}
