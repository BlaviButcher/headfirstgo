package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter file location: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	amount := 0
	sum := 0.0
	for scanner.Scan() {
		amount++
		input := scanner.Text()
		distance, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += distance
	}
	fmt.Printf("Average: %.2f\n", sum/float64(amount))
}
