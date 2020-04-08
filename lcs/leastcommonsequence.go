package main

import (
	"fmt"
	"strings"
)

func main(){
fmt.Println(longestCommonSequence("AGGTAB","GXTXAYB", len("AGGTAB"),len("GXTXAYB")))
}
func longestCommonSequence(str1,str2 string,l1,l2 int)int{
	//mnMatrix := make([][]int{})
	mnMatrix := [7][8]int{}
	str1s := strings.Split(str1,"")
	str2s := strings.Split(str2,"")
	for i:=0;i<=l1;i++{
		for j:=0;j<=l2;j++{
			if i==0 || j==0 {
				mnMatrix[i][j] = 0
			} else if str1s[i-1]==str2s[j-1]{
					mnMatrix[i][j]=1+mnMatrix[i-1][j-1]
				}else{
					mnMatrix[i][j] = max(mnMatrix[i-1][j],mnMatrix[i][j-1])
				}


		}
	}
	return mnMatrix[l1][l2]
}

func max(n1,n2 int) int {
	if n1>n2{
		return  n1
	}
	return n2
}

