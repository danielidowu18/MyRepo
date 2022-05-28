package main

import ("fmt"
		"math/rand"
		"time"
	)
func main (){
	age := rand.Intn(20)
	rand.Seed(time.Now().UnixNano())
    if age >= 18 {
		fmt.Println("You can drink alcohol, but remember to drink responsibly")
	} else {
		fmt.Println("Stay away from alcohol")
	}
	isBankRobberyOn := true
	if (isBankRobberyOn == false) {
		fmt.Println("Sorry, better luck next time")
	} else {
		fmt.Println("Great, remember to be careful of the guards")
	}
	answer := "spider"
	if answer == "chimpanzee" {
		fmt.Println("It's a mammal.")
	} else if answer == "spider" {
		fmt.Println("It's an insect.")
	} else {
		fmt.Println("Then it's unknown")
	}

} 