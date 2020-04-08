package main

import "fmt"

func main(){
	data := []int{7,5,2,23,43,4}

	for i:=1;i< len(data);i++{


		for j:=i;j>0 &&(data[j]>data[j-1]);j--{
			data[j],data[j-1]=data[j-1],data[j]
		}


	}

	fmt.Println(data)
}
