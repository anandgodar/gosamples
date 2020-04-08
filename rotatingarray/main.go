package main

import "fmt"

func main(){
	a:= []int{1,2,3,4,5}
	fmt.Println(rotateR(a,2,len(a)))
	fmt.Println(rotate(a,2,len(a)))
}

func rotate(ar []int,d,n int) []int{
	var newArray []int
	for i:=0;i<d;i++{
		newArray = ar[1:n]
		newArray = append(newArray,ar[0])
		ar = newArray
	}

	return ar
}
func rotateR(ar []int,d,n int) []int{


	ar = append(ar[d:n],ar[0:d]...)
	return  ar
}