package main

import (
	
	"github.com/blavibutcher/headfirst/magazine"
)


func main() {
	subscriber1 := magazine.DefaultSubscriber("Harry")
	subscriber2 := magazine.DefaultSubscriber("Liam")
	magazine.PrintInfo(subscriber1)
	
	magazine.PrintInfo(subscriber2)
	magazine.ApplyDiscount(subscriber2)
	magazine.PrintInfo(subscriber2)
}

