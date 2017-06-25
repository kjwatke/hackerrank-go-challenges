package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Store number of hurdles, max jumping height, and number of drinks needed.
	// var numOfHurdles int
	// var maxHeight int
	// var drinksNeeded int
	var numOfHurdles, maxHeight, drinksNeeded int

	io := bufio.NewReader(os.Stdin)

	// Store values read from Stdin.
	fmt.Fscanln(io, &numOfHurdles, &maxHeight)

	// Store all hurdle heights needed to complete level.
	hurdleHeights := make([]int, numOfHurdles)
	for i := range hurdleHeights {
		fmt.Fscan(io, &hurdleHeights[i])
	}

	// Iterate over hurdleHeights and find the tallest hurdle.
	maxHurdle := 0
	for _, v := range hurdleHeights {
		if v > maxHurdle {
			maxHurdle = v
		}
	}

	if maxHeight > maxHurdle {
		drinksNeeded = 0
	} else {
		drinksNeeded = maxHurdle - maxHeight
	}

	// Print out number of drinks needed.
	fmt.Println(drinksNeeded)
}
