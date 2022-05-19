package main

import (
	"fmt"
	"sort"
)

func main() {

	var number = []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println(number)
	//var maxcountovent int = findMaxConsecutiveOnes(number)
	//maxcountovent = findNumbers(number)
	//fmt.Println(maxcountovent)
	//fmt.Println(sortedSquares(number))
	fmt.Println(duplicateZeros(number))
}

func findMaxConsecutiveOnes(nums []int) int {

	var countmax int = 0
	var memMaxCon int = 0

	for i := 0; i <= len(nums)-1; i++ {

		if nums[i] == 1 {
			countmax++
		} else {

			if countmax > memMaxCon {
				memMaxCon = countmax
				countmax = 0
			}

		}

	}

	if countmax > memMaxCon {
		memMaxCon = countmax
	}

	return memMaxCon
}

func findNumbers(nums []int) int {
	var countmax int = 0

	for i := 0; i < len(nums)-1; i++ {

		if checkDigit(nums[i])%2 == 0 {
			countmax++
		}
	}

	return countmax
}

func checkDigit(number int) int {
	var digit = 0
	for number > 0 {
		number = number / 10
		digit++
	}

	return digit
}

func sortedSquares(nums []int) []int {

	var squares = []int{}

	i := 0
	for i <= len(nums)-1 {
		squares = append(squares, nums[i]*nums[i])
		i++
	}

	sort.Ints(squares)

	return squares

}

func duplicateZeros(arr []int) []int {

	for i := 0; i <= len(arr)-1; i++ {

		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}

	return arr
}
