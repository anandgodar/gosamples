package main

import (
	"fmt"
	"time"
)

func main(){
	in := make(chan int)
	out := make(chan int)
	go producer(in,100*time.Millisecond)
	go producer(in, 150*time.Millisecond)
	go reader(out)
	for i:= range in {
		fmt.Println("Channel",in)
		fmt.Println("value of i :",i)

	}


}


func producer(in chan int, d time.Duration){

	//var i int
	for i:=0;i<100;i++ {

		in<-i
		//i++
		time.Sleep(d)

	}


}

func reader(out chan int ){
	for x:= range out {
		fmt.Println(x)
	}

}