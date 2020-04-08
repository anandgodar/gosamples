package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(longestPalindromSubstr("rmaaamtp"))
}
func longestPalindromSubstr(str string)(ln int, pStr string){
	strArr := strings.Split(str,"")
	s := len(strArr)
	mnMatrix := [8][8]bool{}
	// Allocating zero values to all the matrix value
	for i:=0;i<s;i++{
		for j:=0;j<s;j++{
			mnMatrix[i][j]=false
		}
	}

   for i:=0;i<s;i++{
   		mnMatrix[i][i]=true
   }

   // Check substring of length 2
   var start int = 0
   //var maxLength int = 0
   for i:=0;i<s-1;i++{
   		if str[i]==str[i+1]{
   			mnMatrix[i][i+1] = true
   			start = i
   			ln = 2
      }
   }
   fmt.Println(start)
// Checking substring >2
   for k:=3;k<=s;k++{
   		for i:=0;i<s-k+1;i++{
   			j:= i+k-1

   			if mnMatrix[i+1][j-1] && str[i]==str[j] {
				mnMatrix[i][j] = true

				if k > ln {
					ln = k
					start = i

				}
			}

		}
   }


	return ln,str[start:start+ln]
}
