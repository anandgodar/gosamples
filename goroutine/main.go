package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	for i:=0;i<10;i++{
		go fmt.Printf("Loop at %d\n",i)
	}
	fmt.Println(" Loop Endend")
	time.Sleep(100*time.Millisecond)

	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go printEven(i,&wg)

	}

	wg.Wait()
	val := 0
	for i:=0;i<1000;i++{
		wg.Add(1)
		go testGoroutine(&wg,&val)
	}
	wg.Wait()
	fmt.Printf("The sum is %d \n",val)

}


func printEven(n int , wg *sync.WaitGroup){
	if n%2==0{
		fmt.Printf("%d is Even\n",n)
	}
	defer wg.Done()
}

func testGoroutine(wg *sync.WaitGroup,ptr *int){
	i:= *ptr
	*ptr = i+1
	wg.Done()
}