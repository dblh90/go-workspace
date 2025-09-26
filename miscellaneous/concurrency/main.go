package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatInt(32, 2))
	fmt.Println(Solution(32))
}

func Solution(N int) int {
	// Implement your solution here
	maxBinaryGabCount := 0
	binaryGabCountAtHand := 0
	for N > 0 {
		remainder := N % 2
		N /= 2
		if N <= 0 {
			binaryGabCountAtHand = 0
		}

		if remainder == 0 {
			binaryGabCountAtHand++
		} else {
			if binaryGabCountAtHand > maxBinaryGabCount {
				maxBinaryGabCount = binaryGabCountAtHand
			}
			binaryGabCountAtHand = 0
		}
	}
	return maxBinaryGabCount
}
