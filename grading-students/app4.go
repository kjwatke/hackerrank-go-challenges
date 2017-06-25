// Hackerrank Challenge: https://www.hackerrank.com/challenges/grading

// Take in a single integer to denote number of grades.

// Take in an unknown number of elements as grades, one per line.

// If a grade is within 3 of the next multiple of 5, add the appropriate value to
// round it up to that next multiple, otherwise leave it alone.

// Print out each value of the grades on their own line.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	var numOfGrades int

	// Get the number of grades to review.
	fmt.Fscan(io, &numOfGrades)

	// Create an array to store grades in.
	grades := make([]int, numOfGrades)

	// Get the grades and store them in an array.
	for i := 0; i < numOfGrades; i++ {
		fmt.Fscan(io, &grades[i])
	}

	// Iterate over grades, if within <= 3 of a multiple of 5, round up to multiple
	for i, v := range grades {

		// If the current value is within 3 of the next multiple of 5, add the
		// appropriate value to that element to make it equal to the next multiple.
		if v%5 > 2 && v >= 38 {
			switch v % 5 {
			case 4:
				grades[i]++
			case 3:
				grades[i] += 2
			case 2:
				grades[i] += 3
			default:
			}
		}

		// Print out the value of the current element on it's own line.
		fmt.Println(grades[i])
	}
}
