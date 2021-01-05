package main


import (
	"fmt"
	"errors"
	"log"
)

// Date holds a date 
type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date) SetYear(year int) error{
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}


func main() {
	date := Date{}
	err := date.SetYear(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}