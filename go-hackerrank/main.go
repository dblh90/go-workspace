package main

import "fmt"

func main() {

	arr1 := []int32{2, 6}
	arr2 := []int32{24, 36}

	numsInBetween := getTotalX(arr1, arr2)

	fmt.Println("numsInBetween ", numsInBetween)
}

func getTotalX(a []int32, b []int32) int32 {
	isMoreThan100A := len(a) > 100
	isMoreThan100B := len(b) > 100

	if isMoreThan100A || isMoreThan100B {
		return 0
	}

	var numsInBetween []int32
	var numInQuestion int32 = 1

	chann1 := make(chan bool)
	chann2 := make(chan bool)

	for {
		if numInQuestion == 8 {
			fmt.Println()
		}

		areArrayElemsAllFactorOfNum := false
		isNumFactorOfAllArrayElems := false

		go checkIfArrayElemsAreAllFactorsOfNum(numInQuestion, a, chann1)
		go checkIfNumIsFactorOfAllArrayElems(numInQuestion, b, chann2)

		select {
		case b1 := <-chann1:
			areArrayElemsAllFactorOfNum = b1
		}

		select {
		case b2 := <-chann2:
			isNumFactorOfAllArrayElems = b2
		}

		if areArrayElemsAllFactorOfNum && isNumFactorOfAllArrayElems {
			numsInBetween = append(numsInBetween, numInQuestion)
		}

		if numInQuestion > 100 {
			break
		}
		numInQuestion++
	}

	return int32(len(numsInBetween))
}

func checkIfArrayElemsAreAllFactorsOfNum(num int32, arr []int32, ch chan bool) {
	isFactorForAll := true
	for index := 0; index < len(arr); index++ {
		arrNum := arr[index]

		if arrNum == 0 {
			continue
		}

		if num%arrNum != 0 {
			isFactorForAll = false
			break
		}
	}
	ch <- isFactorForAll
}

func checkIfNumIsFactorOfAllArrayElems(num int32, arr []int32, ch chan bool) {
	isFactorForAll := true
	for index := 0; index < len(arr); index++ {
		arrNum := arr[index]

		if arrNum == 0 {
			continue
		}

		if arrNum%num != 0 {
			isFactorForAll = false
			break
		}
	}

	ch <- isFactorForAll
}
