package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// left = 0 , 97 = -1
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
		fmt.Println("location si",location[s[i]],"left",left,"si",s[i],"maxLen",maxLen)
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	start := -1
	lenMax := 0
	dict := [256]int{}
	for i := range dict {
		dict[i] = -1
	}
	fmt.Println(dict)
	for i, r := range s {
		fmt.Println(i,r)
		if v := dict[r]; v > start {
			start = v
		}
		length := i - start
		fmt.Println("leng",length)
		if length > lenMax {
			lenMax = length
		}
		dict[r] = i
	}
	return lenMax
}

func main(){
	fmt.Println(lengthOfLongestSubstring("abcdeeapqrsruvwxyz"))
	fmt.Println(lengthOfLongestSubstring2("abcdeeapqrsruvwxyz"))
	// a = 97, initiall -1 , location = 1
	// b = 98 left =0  maxlent = 2
	// cd = 99 left=0 max = 3
	// d = 100 left =0 max 4
	// e = 101 left= 0 max 5
	// e = 101 left = 4+1 = 5
	// a = 97
}
