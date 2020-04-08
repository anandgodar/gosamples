package main

import "fmt"

func main(){
	data := []int{7,5,2,23,43,4}

	for i:=0;i< len(data);i++{
		for j:=i+1;j<len(data);j++{
			if data[i]>data[j]{
				c := data[i]
				data[i] = data[j]
				data[j]=c
			}
			fmt.Println("--",data)
		}
		fmt.Println(data)
	}

	fmt.Println(data)
}
