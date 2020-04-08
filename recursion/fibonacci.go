package main

import "fmt"
var fibDict [15]int
func main(){
	// 0 1 1 2
	//i:=10
	t1 :=0
	t2:=1
	sum:=0
	i:=0
	for i<10{
		sum = t1 +t2
		fmt.Print(t1," ")
		t1 =t2
		t2= sum
		i++
	}
	fmt.Println("------Reucsion -----")
	for i:=1;i<=10;i++{
		fmt.Print(fib(i)," ")
	}
}

func  fib (n int )  int {

	if n==0{
		fmt.Println("invalid")
	}
	if n==1{
		return  0
	} else if n==2 {
		return 1
	}

	if fibDict[n]!=0{
		//fmt.Println("iinsed")
		return fibDict[n]
	}else {

		value := fib(n-1) + fib(n-2)
		fibDict[n] = value
		return value
	}
}
/*

 */

