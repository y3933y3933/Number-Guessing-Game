package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

type Difficulty struct {
	Name   string
	Chance int
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	randomNum := rand.IntN(100) + 1
	guessRound := 0
	choices := map[string]Difficulty{
		"1": {
			Name:   "Easy",
			Chance: 10,
		},
		"2": {
			Name:   "Medium",
			Chance: 5,
		},
		"3": {
			Name:   "Hard",
			Chance: 3,
		},
	}

	fmt.Print(`Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: `)

	var difficulty Difficulty

	for {
		reader.Scan()
		choice, exist := choices[reader.Text()]
		if !exist {
			fmt.Println("\nPlease select the difficulty level.\n\nEnter your choice: ")
			continue
		}
		difficulty = choice
		break
	}

	fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game!\n", difficulty.Name)

	for guessRound < difficulty.Chance {
		guessRound++
		fmt.Print("Enter your guess: ")
		reader.Scan()

		num, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println("Incorrect! You should enter a integer.")
			continue
		}
		if num == randomNum {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", guessRound)
			break
		}

		if num > randomNum {
			fmt.Printf("Incorrect! The number is less than %d.\n", num)
			continue
		} else {
			fmt.Printf("Incorrect! The number is greater than %d.\n", num)
		}

	}

}
