package main

import (
	. "fmt"
	"math/rand"
)

var n, score int

func main() {
	for {
		Print("\nPC set the number...")
		a := rand.Intn(6)
		Print("\nGuess the number (0 to 6 ):")
		_, err := Scan(&n)
		if err != nil {
			Print("Error: ", err)
		}

		if n == a {
			score++
			Print("You guessed the number")
			break
		} else {
			Print("You not guessed number,try again")
			for {
				Print("\nAgain: ")
				_, err := Scan(&n)
				if err != nil {
					Print("Error: ", err)
				}

				if a == n {
					score++
					Printf("Great,you guessed the number.Game over.You score is %d", score)
					return
				} else {
					Print("Incorrect. Try again.")
				}
			}
		}

	}
}
