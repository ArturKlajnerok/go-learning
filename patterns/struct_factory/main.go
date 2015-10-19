package main

import (
	"./appliances"
	"fmt"
)

func main() {
	//Request the user to enter the appliance type
	fmt.Println("Enter preferred appliance type")

	//use fmt.scan to retrieve the user's input
	var myType int
	fmt.Scan(&myType)

	//Use the class factory to create an appliance of the requested type
	myAppliance, err := appliances.CreateAppliance(myType)

	//if no errors start the appliance then print it's purpose
	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		//if error encountered, print the error
		fmt.Println(err)
	}

}
