package main

import "fmt"

func main(){
	fmt.Println(knapsack([]int{1,2,5,6},[]int{2,3,4,5},8))
}

func knapsack(profit,weight []int, w int)int {
	v := [5][9]int{}
	for i:=0;i<=len(weight);i++{
		for j:=0;j<=w;j++{
			 if i==0 || j==0 {
			 	v[i][j]=0
			 	continue
			 }

			 if j-weight[i-1]>=0{
				 v[i][j] = max(v[i-1][j],v[i-1][j-weight[i-1]]+profit[i-1])
			 }else{
			 	v[i][j] = v[i-1][j]
			 }
		}
	}
	for  _,va:= range v {
		fmt.Println(va)
	}
	return  v[len(weight)][w]
}

func max(n1,n2 int)int {
	if n1>n2{
		return n1

	}
	return n2
}