package main

import (
	"fmt"
	"sort"
)

func main() {

	var number = []int{3,1,2,4}
	fmt.Println(number)
	//var m = 3
	//var n = 3
	//go run leetcode.govar number2 = []int{2, 5, 6}
	//fmt.Println(number)
	//var maxcountovent int = findMaxConsecutiveOnes(number)
	//maxcountovent = findNumbers(number)
	//fmt.Println(maxcountovent)
	//fmt.Println(sortedSquares(number))
	//fmt.Println(duplicateZeros(number))
	//merge(number, m, number2, n)
	//fmt.Println(checkIfExist(number))
	fmt.Println(sortArrayByParity(number))

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
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[0:n]...)
	sort.Ints(nums1)

	fmt.Println(nums1)
}

func removeElement(nums []int, val int) int {

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == val {

			nums = append(nums[:i], nums[i+1:len(nums)]...)

			i = i - 1

		}

	}
	return len(nums)
}
func removeDuplicates(nums []int) int {

	sort.Ints(nums)

	for i := 0; i <= len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {

			nums = append(nums[:i], nums[i+1:len(nums)]...)

			i = i - 1
		}

	}
	return len(nums)
}
func checkIfExist(arr []int) bool {

	var rs = false
	for i := 0; i < len(arr)-1; i++ {
		rs = linearserch(i, arr)
		if rs {
			break
		}
	}
	return rs

}
func linearserch(item int, arr []int) bool {
	var found bool = false
	for i := 0; i <= len(arr)-1; i++ {
		if item != i && (arr[i]*2) == arr[item] {
			found = true
			break
		}
	}
	return found
}
func validMountainArray(arr []int) bool {
    
	var increase = false;
	var decrease = false;

	if len(arr)< 3 {
		return false
	}

	if arr[1]<arr[0]{
		return false
	}
		for i:=1 ;i<=len(arr)-1;i++{
			if (arr[i] == arr[i - 1]) {
					return false
				}
			if (arr[i] >arr[i - 1]) {
				
				if (decrease == true) {
					return false
				}
				increase = true
			}
			if (arr[i] < arr[i - 1]) {
				decrease = true
			}


		}

	if (increase == true && decrease == true) {
		return true
	} else {
		return false
		}
	return false
}

func replaceElements(arr []int) []int {
    
	
	var maxnum int = arr[len(arr)-1]

	for i:=len(arr)-2;i>=0;i--{
		
		var mem int = arr[i]
		arr[i]= maxnum
		if(mem>maxnum) {
			maxnum = mem
		}		
	}

	arr[len(arr)-1] = -1

	return arr
}
func moveZeroes(nums []int) []int {
    
	if len(nums)==0{
		return nums
	}
	
	var newarr  = []int{}
	var cntZero  int = 0

	for i:=0;i<=len(nums)-1;i++{
		
		if nums[i]!=0{
			newarr = append(newarr, nums[i])
		}else{
			cntZero++
		}

		
	}

	for i:=0;i<cntZero;i++{
		newarr = append(newarr,0 )
	}

	for i:=0;i<=len(newarr)-1;i++{
		
	nums[i]=newarr[i]
		
	}
	
	return nums
}

func sortArrayByParity(nums []int) []int {
    
    var  evearr  = []int{}
	var  oddarr  = []int{}

	for i:=0;i<=len(nums)-1;i++{
		
		if nums[i]%2==0{
			evearr = append(evearr, nums[i])
		}else{
			oddarr = append(oddarr, nums[i])
		}

		
	}

	evearr = append(evearr, oddarr...) 

	return  evearr
}

