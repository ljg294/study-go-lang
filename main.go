package main

import "fmt"

func main() {
	fmt.Println("This is the sample project for studying go-lang")

	var sampleText = "sample text with go-lang"
	fmt.Println("this is a", sampleText)

	var myAge = 30
	myAge++
	fmt.Printf("I am %v years old\n", myAge)

	const thisYear = 2023
	fmt.Printf("It is %v now\n", thisYear)

	var userName string
	var userHeight int
	var isMale bool

	userName = "Jay"
	userHeight = 180
	isMale = true
	fmt.Printf("my name is %v \n", userName)
	fmt.Printf("i am %v high \n", userHeight)
	fmt.Printf("I am male, %v \n", isMale)

	fmt.Printf("userName = %T userHeight = %T isMale = %T \n", userName, userHeight, isMale)

	fmt.Printf("=========================== \n")

	var number int
	number = 940621
	fmt.Printf("number is %v \n", number)
	fmt.Printf("number is at %v \n", &number)

	var emailAddress string
	fmt.Println("type your email address : \n")
	fmt.Scan(&emailAddress)
	fmt.Printf("Is your email address %v \n ?", emailAddress)
}
