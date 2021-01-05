package main

import (
	"fmt"
	"github.com/blavibutcher/headfirst/calender"
	"log"
)

func main() {
	date := calander.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}