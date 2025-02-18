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

	firstInputAsString := strings.Split(firstInput, " ")
	firstInputAsInt := make([]int32, 0)

	for _, value := range firstInputAsString {
		convertedNum, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting the first input to int", value, err)
			return
		}
		firstInputAsInt = append(firstInputAsInt, int32(convertedNum))
	}

	fmt.Println(pickingNumbers(firstInputAsInt))

}

func pickingNumbers(a []int32) int32 {
	// Write your code here
	max := int32(0)
	return max
}
