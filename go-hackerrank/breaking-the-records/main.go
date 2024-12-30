package main

import "fmt"

func main() {

	numOfRecordCount := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Println("numOfRecordCount ", numOfRecordCount)

}

func breakingRecords(scores []int32) []int32 {
	// Write your code here

	var brokenRecordsCount = []int32{0, 0}

	highestRecordYet := int32(0)
	lowestRecordYet := int32(0)

	for index, value := range scores {
		if index == 0 {
			highestRecordYet = value
			lowestRecordYet = value
			continue
		}

		if value > highestRecordYet {
			brokenRecordsCount[0] = brokenRecordsCount[0] + 1
			highestRecordYet = value
		}

		if value < lowestRecordYet {
			brokenRecordsCount[1] = brokenRecordsCount[1] + 1
			lowestRecordYet = value
		}
	}

	return brokenRecordsCount
}
