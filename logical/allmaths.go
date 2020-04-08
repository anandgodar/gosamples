package main

import "fmt"

func main(){
	FindConsecutiveSum(20)
	arr:=[]int{5,6}
	fmt.Println(CalculateTheMinTimeToProduceMItemsFromMachines(arr,2,11))
}
//Consecutive sum is the sum of continous number lets say if we give 10 , then answer will be 1,2,3,4
// 15 will be 1,2,3,4,5
// 20 will be 2,3,4,5,6
func FindConsecutiveSum(n int){
	start:=1
	end := (n+1)/2
	for start<=end{
		sum := 0
		for i := start; i <=end; i++ {
			sum = sum +i
			if sum == n {
				for j:=start;j<=i;j++{
					fmt.Println(j)

				}
				break;
			}
		}
		start++;
		sum=0

	}
}

//Consider the machine which produce item at a given sec , find the number of sec to produce the m items
// Brute force method and binary search method
// Implementation 1 by brute force


// Return the minimum time required to
// produce m items with given machines.
// m is the items to produce, n is the number of ite
func CalculateTheMinTimeToProduceMItemsFromMachines(timesForEachMachine []int,n,m int) int {
	//Initialize time to 0
	t:=0
	for {
		items :=0
		for i:=0;i<n;i++{
			items += t/timesForEachMachine[i]
		}
		// if number of items == m , then return the time required to produce m items from the given n number of machines
		if items>=m{
			return t
		}
		t++
	}
}


