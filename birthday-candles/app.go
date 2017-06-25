package main

import (
	"bufio"
	"fmt"
	"os"
)

// Given x number of candles, and you can only blow out the tallest
// candles in a cake, return the number of candles that can be blown out.
// Ex: given 4 candles, and heights of [1, 2, 3, 4], an int of 1 should be
// returned since only 4 can be blown out.
// Inputs: 1st line of Stdin is the number of candles.
// 2nd line of Stdin is a space seperated group of lengths.
// Output: how many of only the tallest candles can be blown out.
func main() {
	var candles int
	var max int
	var count int

	// Make io to read from Stdin
	io := bufio.NewReader(os.Stdin)

	// Read the number of candles from Stdin.
	fmt.Fscan(io, &candles)

	// Create an array of ints the size of candles.
	lengths := make([]int, candles)

	// Loop over Stdin and store each length in an array.
	for i := 0; i < candles; i++ {
		fmt.Fscan(io, &lengths[i])
	}

	// Find the max value in lengths[]
	max = lengths[0]
	for i, v := range lengths {
		if v > max {
			max = lengths[i]
		}
	}

	// Found max, now find how many occurances are found in lengths[]
	for _, v := range lengths {
		if v == max {
			count++
		}
	}

	// Print the number of candles that can be blown out.
	fmt.Println(count)
}
