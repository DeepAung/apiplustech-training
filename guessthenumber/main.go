package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	cap := 0
	fmt.Print("Hi, this is a number guessing game. Please enter the upper bound: ")
	fmt.Scan(&cap)

	goal := randInt(cap)
	guess := 0

	maximumAttempts := int(math.Ceil(math.Log2(float64(cap))))

	// for try := 0; try < maximumAttempts; try++ {
	for range maximumAttempts {
		fmt.Print("Please Enter your guess: ")
		fmt.Scan(&guess)
		switch {
		case guess == goal:
			fmt.Println("Yay!, you got it right.")
			return
		case guess > goal:
			fmt.Println("Nay, your guess is too large.")
		default:
			fmt.Println("Nah, your guess is too small.")
		}
	}
	fmt.Println("You ran out of attempts, Game Over.")
}

// This function returns random integer from 0 to cap
func randInt(cap int) int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randInstance := rand.New(source)
	goal := randInstance.Intn(cap + 1)
	return goal
}
