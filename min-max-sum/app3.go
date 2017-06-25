package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMaxMin(arr [5]int64) (int64, int64) {
	var max = arr[0]
	var min = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return max, min
}

func main() {
	var nums [5]int64
	var sum int64
	var sums [5]int64
	io := bufio.NewReader(os.Stdin)

	// Store 5 integers needed to calculate min and max.
	for i := 0; i < 5; i++ {
		fmt.Fscan(io, &nums[i])
	}

	// Outer loop is used to skip i on each iteration and reset sum on each pass.
	for i := 0; i < 5; i++ {
		sum = 0

		// Inner loop is used to sum up the 4 of 5 values and store them in array.
		for j := 0; j < 5; j++ {
			if i == j {
				continue
			}

			sum += nums[j]
		}

		sums[i] = sum
	}

	max, min := findMaxMin(sums)
	fmt.Println(min, max)
}
