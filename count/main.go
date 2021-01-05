package main

import (
	"github.com/blavibutcher/datafile"
	"log"
	"fmt"
	"sort"
)

func main() {
	names, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	votes := make(map[string]int)
	for _, name := range names {
		votes[name]++
	}
	// This will print a randomly list each time	
	// for key, value := range votes {
	// 	fmt.Printf("%s: %d\n", key, value)
	// }


	//order fix is rest of this
	// order fix
	ordered := make([]string, 0)

	for name := range votes {
		ordered = append(ordered, name)
	}

	sort.Strings(ordered)
	for _, name := range ordered {
		fmt.Printf("%s: %d\n", name, votes[name])
	}
	
	
}

