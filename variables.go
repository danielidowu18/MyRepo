package main

import "fmt"

var myName, title, email, phoneNumber string
var experienceInWeeks, myAge int
func main() {
	email = "danielakinyeluwa@gmail.com"
	phoneNumber = "09133476365"
	myAge = 16
	myName = "Daniel Idowu"
	title = "backend developer"
	experienceInWeeks = 4
	fmt.Println("THIS IS A LITTLE INFO ABOUT ME")
	fmt.Println("My name is", myName)
	fmt.Println("I am a", title)
	fmt.Println("I have been working for", experienceInWeeks, "weeks")
	fmt.Println("You can reach me on ", email, "or", phoneNumber)

}
