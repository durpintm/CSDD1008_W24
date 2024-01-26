package main

import (
	"fmt"
	"math/rand"
)

//Method to Generate Random number from a range of min and max number
func GenerateRandomNumber() int {
	//Declaring variables to store minimum number for range
	var min, max int

	//Ask for user input
	fmt.Println("Enter minimum and maximum range number: ")

	//store minimum and maximum number in variable min and max
	fmt.Scanln(&min, &max)

	//Check if max number is less or equal to min number
	if(max <= min){
		fmt.Println("Invalid Numbers: Please enter valid min and max numbers")
		return -1
	}
	//Generate random number between max and min numbers
	randomNumber := rand.Intn(max - min) + min

	//return random number to the caller
	return randomNumber
}