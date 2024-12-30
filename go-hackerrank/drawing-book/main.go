package main

import "fmt"

func main() {

	numOfPageTurns := pageCount(37455, 29835)
	fmt.Println("numOfPageTurns ", numOfPageTurns)

}

func pageCount(n int32, p int32) int32 {
	// Write your code here
	pageInTheMiddle := n / 2
	var index int32
	var numOfTurns int32

	if p == 1 || p == n {
		return 0
	}

	if n-p == 1 && n%2 == 0 {
		return 1
	}

	if n%2 == 1 && p%2 == 0 {
		n = n - 1
	}

	if p <= pageInTheMiddle {
		// normal for loop
		for index = 0; index < n; {
			if index != p && index+1 != p {
				numOfTurns++
				index = index + 2
			} else {
				break
			}
		}

	} else {
		// reverse for loop

		for index = n; index >= 0; {
			if index != p && index-1 != p {
				numOfTurns++
				index = index - 2
			} else {
				break
			}
		}
	}

	return numOfTurns
}
