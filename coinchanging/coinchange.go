package main

import "fmt"

func main(){
	fmt.Println(coinChange([]int{1,2,5},11))
}
func coinChange(coins []int, t int) int{

	combinations := make([]int, t+1)
	combinations[0] = 0
	for i:=1;i<=t;i++{
		combinations[i]=t+1
	}
	for i:=1;i< len(combinations);i++{
		for j:=0;j<len(coins);j++{

			if i>=coins[j]{
				combinations[i] = min(combinations[i],combinations[i-coins[j]]+1)
			}

		}
	}
	fmt.Println(combinations)
	return  combinations[11]
}

func min(n1,n2 int)int {
	if n1>n2{
		return n2
	}
	return n1
}