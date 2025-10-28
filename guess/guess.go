package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	/*
	low := 1
	high := 100
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please think of a number between ", low, "and ",high)
	fmt.Println("Enter when ready")
	scanner.Scan()

	tries := 0
	// infinite loop
	for {
		// condition to check i a user is lying
		if low > high {
			fmt.Println("user is lying")
		}

		guess := (low + high)/2
		fmt.Println("i guess the number is ",guess)
		fmt.Println("Is that? :")
		fmt.Println("a. to high")
		fmt.Println("b. to low")
		fmt.Println("c. correct")

		scanner.Scan()
		response := scanner.Text()
		if response == "a" {
			high = guess - 1
			tries++
		} else if response == "b" {
			low = guess + 1
			tries++
		} else if response == "c" {
			fmt.Println("correct guess")
			fmt.Println("total tries till correct: ",tries)
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
	*/

	low := 1
	high := 100
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("guess a number between 1 and 100")
	fmt.Println("enter to continue")
	scanner.Scan()
	tries := 0

	for {
		guess := (low + high)/2
		if low > high {
			fmt.Println("user is lying")
		}
		fmt.Println("the guess is",guess)
		fmt.Println("choices you have")
		fmt.Println("a : to high")
		fmt.Println("b : to low")
		fmt.Println("c : correct")
		scanner.Scan()
		result := scanner.Text()

		if result == "a" {
			high = guess - 1
			tries++
		} else if result == "b" {
			low = guess + 1
			tries++
		} else if result == "c" {
			fmt.Println("correct")
			fmt.Println("tries:",tries)
			break
		} else {
			fmt.Println("invalid input")
		}
	}
}


