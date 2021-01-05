package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please enter a width: ")
	reader := bufio.NewReader(os.Stdin)
	width, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("Please enter a height: ")
	height, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	width = strings.TrimSpace(width)
	height = strings.TrimSpace(height)

	iWidth, err := strconv.ParseFloat(width, 64)

	if err != nil {
		log.Fatal(err)
	}

	iHeight, err := strconv.ParseFloat(height, 64)

	if err != nil {
		log.Fatal(err)
	}

	area, err := area(iWidth, iHeight)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nThe area of paint need is %.2f meters squared\n", area)

	cans := area / 10
	cans = math.Ceil(cans)
	fmt.Printf("You need %.0f can/s of paint\n", cans)

}

func area(width, height float64) (float64, error) {
	// var err error
	// if width < 0 {
	// 	err = fmt.Errorf("Width is negative (%f), this is not allowed", width)
	// } else if height < 0 {
	// 	err = fmt.Errorf("Height is negative (%f), this is not allowed", height)
	// }
	// return width * height, err

	if width < 0 {
		return 0, fmt.Errorf("Width is negative (%f), this is not allowed", width)
	} else if height < 0 {
		return 0, fmt.Errorf("Height is negative (%f), this is not allowed", height)
	}
	return width * height, nil
}
