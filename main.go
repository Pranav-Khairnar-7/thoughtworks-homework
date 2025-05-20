package main

import (
	"fmt"
	user "pranav/day1"
	concurrency "pranav/day2"
	// concurrency "pranav/day2" required for checking day2 homework
	// user "pranav/day1" required for checking day1 homework
)

func main() {

	// Uncomment the function you want to test !!

	//Homework day 1

	//case 1 positve an it will work
	name := "Pranav"
	age := 26
	email := "abc@gmail.com"

	//case 2 this will throw error for email field
	// name := "Pranav"
	// age := 26
	// email := "abcgmail.com"

	user, err := user.NewUser(name, age, email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User created :-", user)
	fmt.Println("Is adult :-", user.IsAdult())

	//Homework day 2
	concurrency.ConcurrentGreeter()
	// concurrency.ConcurrentSumWithChanels()
	// concurrency.ConcurrentSumWithWg()
	// concurrency.PingPongChanels()
}
