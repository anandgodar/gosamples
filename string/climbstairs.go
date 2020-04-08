package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n==1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i:=2;i<n;i++ {
		dp[i] = dp[i-2]+dp[i-1]  // 0,1 d[2] 3 , d[3] = 2+3
	}

	return dp[n-1]
}

func main(){
	fmt.Println(climbStairs(5))

	// 1
	//2+2
	// 1+2+1
	//2+1+1
	//112
}
