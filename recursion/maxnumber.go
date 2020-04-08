package main

import "fmt"

func maxNumber(arr []int)(max int){
	fmt.Println("called")

	// base case
	if len(arr) == 2 {
		fmt.Println("base case")
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}

	// recursive case
	subMax := maxNumber(arr[1:])

	if arr[0] > subMax {
		fmt.Println("arr[0] > subMax")
		return arr[0]
	} else {
		fmt.Println("arr[0] < subMax")
		return subMax
	}
}

func main(){
	arr := []int{3,5,3,2,5,8,23,4}
	//fmt.Println(maxNumber(arr))
	fmt.Println(largestM(arr[1:],arr[0]))
}

func largestM(arr []int, largest int) int {

	if len(arr)>0{
		if arr[0]>largest{
			largest = arr[0]

		}
		largest = largestM(arr[1:],largest)
	}
	return largest
}
