package main

import "fmt"

var arr [1000]int
func main(){
	fmt.Print(countSum(20,[]int{5,3,10},2))
	//printallCombination(0,5)
}

func printallCombination(i,n int){

	if n==0{
		printArray(arr,i)
	}else if n>0{
		for k:=1;i<=3;k++{
			arr[i] = k
			printallCombination(n-k,i+1)
		}
	}
}
func printArray(arr [1000]int,j int){
	for i:=0;i<j;i++{
		fmt.Print(arr[i])
		fmt.Println("-")
	}
}

func countSum(n int, scores []int, end_index int) int {
	if n == 0 {
		return 1
	}else if n < 0 || end_index < 0 {
		return 0
	}else if scores[end_index] > n {
		return countSum(n, scores, end_index-1)
	} else {
		return countSum(n-scores[end_index], scores, end_index) + countSum(n, scores, end_index-1)
	}
}