package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(lcs("CLCL","LCLC"))
}

func lcs(s1,s2 string)(max int,str string){
	str1 := strings.Split(s1,"")
	str2 := strings.Split(s2,"")
	fmt.Println(str1,str2)

	str = ""
	mnMatrix := [4][4]int{}
	for i:=0;i<len(str1);i++{
		for j:=0;j<len(str2);j++{
			if str1[i]==str2[j]{
				if i==0 || j==0 {
					mnMatrix[i][j] = 1
					max = 1
					//str = str1[i]
				}else{
					mnMatrix[i][j] = mnMatrix[i-1][j-1]+1
					max = mnMatrix[i][j]
					str = ""
					for k:=max;k>=1;k--{
						str = str + str2[k]
						//fmt.Println(str)
					}
				}


			}else{
				mnMatrix[i][j] = 0
			}
		}
	}
	fmt.Println(mnMatrix)
	return max, str
}
