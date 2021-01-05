package main

import "github.com/blavibutcher/headfirst/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}