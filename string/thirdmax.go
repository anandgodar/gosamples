package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
ret := []int{math.MinInt64, math.MinInt64, math.MinInt64}

for _,n := range nums{
	if n == ret[0] || n == ret[1] || n == ret[2] {
		continue
	}
	if n > ret[0]{
				ret = []int {n, ret[0], ret[1]}
			}else if n>ret[1] {
				ret = []int{ret[0], n, ret[1]}
			}else if n>ret[2]{
				ret = []int{ret[0], ret[1], n}
		}
	}
	fmt.Println(ret)
	if ret[2] == math.MinInt64{
		return ret[0]
	}else{
		return ret[2]
	}

}

func main(){
	fmt.Println(thirdMax([]int{45,46,50,48,49}))
}