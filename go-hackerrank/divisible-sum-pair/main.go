package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add the expected input")
	fmt.Println("---------------------")

	firstInput, _ := reader.ReadString('\n')
	firstInput = strings.TrimSuffix(firstInput, "\n")
	secondInput, _ := reader.ReadString('\n')
	secondInput = strings.TrimSuffix(secondInput, "\n")

	firstInputAsString := strings.Split(firstInput, " ")
	firstInputAsInt := make([]int32, 0)
	secondInputAsString := strings.Split(secondInput, " ")
	secondInputAsInt := make([]int32, 0)

	for _, value := range firstInputAsString {
		convertedNum, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting the first input to int", value, err)
			return
		}
		firstInputAsInt = append(firstInputAsInt, int32(convertedNum))
	}

	for _, value := range secondInputAsString {
		convertedNum, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting the first input to int", value)
			return
		}
		secondInputAsInt = append(secondInputAsInt, int32(convertedNum))
	}

	numberDivisible := divisibleSumPairs(firstInputAsInt[0], firstInputAsInt[1], secondInputAsInt)
	fmt.Println(numberDivisible)
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {

	count := int32(0)
	for index := 0; index < len(ar); index++ {
		for j := index + 1; j < len(ar); j++ {
			if (ar[index]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}
