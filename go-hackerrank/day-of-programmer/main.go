package main

import (
	"fmt"
	"strconv"
)

func main() {
	year := int32(1800)
	fmt.Println(dayOfProgrammer(year))
}

func dayOfProgrammer(year int32) string {
	isJulianCalendar := year >= 1700 && year <= 1917

	switch isJulianCalendar {

	case true:
		isLeapYear := year%4 == 0
		if isLeapYear {
			return "12.09." + strconv.Itoa(int(year))
		} else {
			return "13.09." + strconv.Itoa(int(year))
		}

	case false:
		if year == 1918 {
			return "26.09.1918"
		}

		isDivisibleBy400 := year%400 == 0
		isDivisibleBy4 := year%4 == 0
		isDivisibleBy100 := year%100 == 0

		if isDivisibleBy400 || (isDivisibleBy4 && !isDivisibleBy100) {
			return "12.09." + strconv.Itoa(int(year))
		} else {
			return "13.09." + strconv.Itoa(int(year))
		}
	default:
		return "Invalid year"
	}
}
