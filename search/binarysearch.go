package main

import "fmt"

func main(){
	search := 13
	searchIn := []int{1,2,3,4,5,6,13}
	index := BinarySearch(search,searchIn)
	fmt.Println(searchIn[index],index)

}

func BinarySearch(searchFor int,searchIn []int) int {
	 var lo int = 0
	 var hi int = len(searchIn)-1

	 for lo<=hi {
		 var midPoint int = lo + (hi-lo)/2
		 var midValue int = searchIn[midPoint]
		 if midValue == searchFor {
			 return midPoint
		 } else if midValue < searchFor {
			 lo = midPoint + 1
		 } else {
			 hi = midPoint - 1
		 }
	 }
	return -1
}