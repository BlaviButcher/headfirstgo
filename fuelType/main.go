package main

import "fmt"

type Gallons float64
type Litres float64
type Millilitres float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Litres(4.5)
	carFuel += Litres(40.0).ToGallons()
	busFuel += Gallons(30.0).ToLitres()
	fmt.Printf("Car Fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus Fuel: %0.1f litres\n", busFuel)
}

func (l Litres) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Millilitres) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLitres() Litres {
	return Litres(g * 3.785)
}