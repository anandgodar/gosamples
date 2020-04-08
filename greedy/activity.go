package main

import "fmt"

/*
Source : https://www.geeksforgeeks.org/activity-selection-problem-greedy-algo-1/
activity selection program
 */

 func activitySelection(start,end []int,n int  )[]int{
 	i:=0
 	output:= make([]int,0)
 	output = append(output,0)
 	for j:=1;j<n;j++{
		if start[j]>=end[i]{
			output = append(output,j)
			i=j
		}
	}

 	return output
 }
func main(){
	fmt.Println(activitySelection([]int{1, 3, 0, 5, 8, 5},[]int{2, 4, 6, 7, 9, 9},6))
}

