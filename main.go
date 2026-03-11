package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println(
		"A random number will be drawn. Try to guess it. The number is an integer between 0 and 100",
	)
	
	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("What is your guess?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Your guess must be an integer")
			return 
		}

		switch {
			case guessInt < x:
				fmt.Println("You're wrong. The drawn number is greater than", guessInt)
			case guessInt > x:
							fmt.Println("You're wrong. The drawn number is less than", guessInt)
			case guessInt == x:
				fmt.Printf(
					"Congratulations! You got it right! The number was: %d\n " +
					"You got it right in %d attempts.\n" +
					"These were your attempts: %v\n",
					x, i+1,  guesses[:i],
				)
				return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"Unfortunately you didn't guess the number, which was: %d\n " +
		"You had 10 attempts.\n" +
		"These were your attempts: %v\n", x, guesses)

}