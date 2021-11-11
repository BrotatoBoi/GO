// ~~ ------------------------------------------------------------------ ~~ \\
// ~~~ ------------- Programmer: AJ Cassell (@BrotatoBoi). -------------- ~~~ \\
// ~~~~ -------------- Program Name: Rock Paper Scissors. ---------------- ~~~~ \\
// ~~~~~ ------------------- Date: November/112/2021. --------------------- ~~~~~ \\
// ~~~~~~ ---- Description: This program is a rock paper scissors game. ---- ~~~~~~ \\
// ~~~~~~~ ------------------------------------------------------------------ ~~~~~~~ \\

// ~ This is the Main File. ~ \\
package main

// ~ Import the packages. ~ \\
import (
	"fmt"
	"math/rand"
	"os"
)

// ~ Get the Computers Choice. ~ \\
func getComputerChoice() string {
	// ~ Output variable. ~ \\
	var computerChoice string

	// ~ Create a random number between 1 and 3. ~ \\
	var randomNumber int = 1 + (rand.Intn(3))

	// ~ Return the computers choice based on the random number. ~ \\
	switch randomNumber {
	case 1:
		computerChoice = "Rock"
	case 2:
		computerChoice = "Paper"
	case 3:
		computerChoice = "Scissors"
	}

	return computerChoice
}

// ~ Play Again? ~ \\
func playAgain() {
	var again string

	fmt.Print("Do you want to play again?: ")
	fmt.Scan(&again)

	if again == "y" || again == "yes" {
		main()
	} else if again == "n" || again == "no" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	} else {
		fmt.Println("Please enter a valid response.")
		playAgain()
	}
}

// ~ The Main Function. ~ \\
func main() {
	// ~ Create Variables. ~ \\
	var userChoice string
	var computerChoice string

	// ~ Welcome the user. ~ \\
	fmt.Println("Welcome to the Rock Paper Scissors Game!")

	// ~ Get the Users Input. ~ \\
	for {
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&userChoice)

		// ~ Check if the user input is valid. ~ \\
		if userChoice == "r" || userChoice == "p" || userChoice == "s" ||
			userChoice == "rock" || userChoice == "paper" || userChoice == "scissors" {
			// ~ Turn shortform into longform. ~ \\
			if userChoice == "r" {
				userChoice = "rock"
			} else if userChoice == "p" {
				userChoice = "paper"
			} else if userChoice == "s" {
				userChoice = "scissors"
			}

			break
		} else {
			fmt.Println("Invalid Input!")
		}
	}

	// ~ Get computers choice. ~ \\
	computerChoice = getComputerChoice()

	// ~ Show the choices. ~ \\
	fmt.Println("You chose: " + userChoice)
	fmt.Println("The computer chose: " + computerChoice)

	// ~ Compare the choices. ~ \\
	if userChoice == computerChoice {
		fmt.Println("It's a tie!")
	} else if userChoice == "rock" && computerChoice == "paper" ||
		userChoice == "paper" && computerChoice == "scissors" ||
		userChoice == "scissors" && computerChoice == "rock" {
		fmt.Println("You lose!")
	} else {
		fmt.Println("You win!")
	}

	// ~ Ask if user wants to play again. ~ \\
	playAgain()
}
