package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	fmt.Println(migratoryBirds(firstInputAsInt))
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	occurrenceCounter := make(map[int32]int32)
	for _, birdType := range arr {
		occurrenceCounter[birdType]++
	}

	pairs := make([][2]interface{}, 0, len(occurrenceCounter))
	for k, v := range occurrenceCounter {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1].(int32) == pairs[j][1].(int32) {
			return pairs[i][0].(int32) < pairs[j][0].(int32)
		}
		return pairs[i][1].(int32) > pairs[j][1].(int32)
	})

	return pairs[0][0].(int32)
}
