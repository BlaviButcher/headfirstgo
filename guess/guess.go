package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a seed to make things more random
	currentTime := time.Now().Unix()
	rand.Seed(currentTime)
	// Get random number
	target := rand.Intn(100) + 1

	fmt.Println("I have chosen a random number between 1 and 100.")

	success := false
	for x := 10; x > 0; x-- {
		fmt.Printf("You have %d guesses remaining!\n", x)
		fmt.Println("What number am I thinking?: ")

		// Get user input
		reader := bufio.NewReader(os.Stdin)
		uInput, err := reader.ReadString('\n')
		fmt.Println("")

		if err != nil {
			log.Fatal(err)
		}

		uInput = strings.TrimSpace(uInput)

		guess, err := strconv.Atoi(uInput)

		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Too High!")
		} else if guess < target {
			log.Println("Too Low!")
		} else {
			log.Println("Correct")
			success = true
			break
		}
		if success == true {
			fmt.Println("YOU WIN!")
		} else {
			fmt.Println("YOU LOSE!")
		}
	}
}
