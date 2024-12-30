package main

import "fmt"

func main() {

	squareVals := []int32{2, 3, 4, 4, 2, 1, 2, 5, 3, 4, 4, 3, 4, 1, 3, 5, 4, 5, 3, 1, 1, 5, 4, 3, 5, 3, 5, 3, 4, 4, 2, 4, 5, 2, 3, 2, 5, 3, 4, 2, 4, 3, 3, 4, 3, 5, 2, 5, 1, 3, 1, 4, 2, 2, 4, 3, 3, 3, 3, 4, 1, 1, 4, 3, 1, 5, 2, 5, 1, 3, 5, 4, 3, 3, 1, 5, 3, 3, 3, 4, 5, 2}
	d := int32(26)
	m := int32(8)

	numOfTimes := birthday(squareVals, d, m)
	fmt.Println("Hello, playground ", numOfTimes)
}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	numOfTimes := int32(0)
	var index int32
	for index = 0; index < int32(len(s)); index++ {
		var currentCount = s[index]
		if currentCount == d && m == 1 {
			numOfTimes++
			continue
		}

		var sum int32
		var count int32 = 1

		if index+1 >= int32(len(s)) {
			break
		}

		subSlice := s[index+1 : len(s)-1]

		for _, val := range subSlice {
			sum += val
			count++
			if count == m {
				break
			}
		}

		if currentCount+sum == d && count == m {
			numOfTimes++
		}
	}

	return numOfTimes
}
