package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	min := 5
	max := 200
	count := 0

	clear()
	fmt.Println("Think of a number between", min, "and", max)
	fmt.Println("Press enter when ready")
	scanner.Scan()

	//*
	// this will search using a random number generator
	rand.Seed(time.Now().Unix())
	guess := min + rand.Intn(max-min+1)
	for {
		fmt.Println("I guess:", guess)
		fmt.Println("\n\n(a) Too high")
		fmt.Println("(b) Too low")
		fmt.Println("(c) Correct!")
		count++
		scanner.Scan()

		resp := scanner.Text()

		if resp == "a" {
			max = guess - 1
			checkSpace(min, max, guess)
			// displayVars(min, max, guess)
			guess = min + rand.Intn(max-min)
		} else if resp == "b" {
			min = guess + 1
			checkSpace(min, max, guess)
			// displayVars(min, max, guess)
			guess = min + rand.Intn(max-min)
		} else if resp == "c" {
			clear()
			fmt.Println("I win!")
			fmt.Println("It took me", count, "guesses.")
			break
		} else {
			fmt.Println("Invalid input, try again")
			displayVars(min, max, guess)
		}
	}
	//*/
	/*
		// This will do a binary search in the available space.
		// This is the most efficient way of doing things
		for {
			guess := (min + max) / 2
			fmt.Println("I guess the number is", guess)
			fmt.Println("Is that...")
			fmt.Println("(a) Too high")
			fmt.Println("(b) Too low")
			fmt.Println("(c) Correct!")
			scanner.Scan()

			resp := scanner.Text()

			if resp == "a" {
				max = guess - 1
				checkSpace(min, max, guess)
				displayVars(min, max, guess)
			} else if resp == "b" {
				min = guess + 1
				checkSpace(min, max, guess)
				displayVars(min, max, guess)
			} else if resp == "c" {
				fmt.Println("I win!")
				break
			} else {
				fmt.Println("Invalid input, try again")
				displayVars(min, max, guess)
			}

		}
		//*/

}

func displayVars(min int, max int, guess int) {
	fmt.Println("Last guess:", guess)
	fmt.Println("New min:", min)
	fmt.Println("New max:", max)
	fmt.Print("\n")
}

func checkSpace(min int, max int, guess int) {
	clear()
	if min > max {
		fmt.Print("\n\n")
		fmt.Println("You lied along the way or chose a number out of range")
		fmt.Println("The the minimum cannot be more than the maximum")
		displayVars(min, max, guess)
		os.Exit(1)
	}
}

func clear() {
	cmd := exec.Command("clear") // macOS and Linux functional
	cmd.Stdout = os.Stdout
	cmd.Run()
}
