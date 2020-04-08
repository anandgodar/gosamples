package main

import "fmt"

func max(a int, b int) int {
	if a>b {
		return a
	}else{
		return b
	}
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0{
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	a := 0
	b := 0
	//maxSum := 0

	for i:=0; i<n; i++ {
		if i%2==0{
			a = a+nums[i]
		}else{
			b = b+nums[i]
		}
		//maxSum = max(a+nums[i], b)
		//a = b
		//b = maxSum
	}

	return max(a,b)
}
func main(){
	fmt.Println(rob1([]int{1,20,34,60,90,120}))
}



func rob1(nums []int) int {
	n := len(nums)
	if n == 0{
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	sums := make([]int, n)
	sums[0] = nums[0]
	sums[1] = max(nums[0], nums[1])

	for i:=2;i<n;i++ {
		sums[i] = max(sums[i-2]+nums[i], sums[i-1])
	}

	return sums[n-1]
}