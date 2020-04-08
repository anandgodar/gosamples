package main

import (
	"fmt"
	"sync"
)

func main(){
	done := make(chan struct{})
	defer close(done)
	ch:= gen(done,2,3,4)


	o1 := produce(done,ch,"one")
	o2 := produce(done,ch,"two")
	//1.
	// fmt.Println(<-ou)
	//fmt.Println(<-ou)

	//2
	// for m:=range ou {
	//	fmt.Println(m)
	//}

	//3
	finalOUt := merge(done,o1, o2)
	//for n := range merge(done,o1, o2) {
		fmt.Println(<-finalOUt)
	//}


}

func gen(done <-chan struct{},nums... int ) <-chan int {
	out := make(chan int)
	go func() {
	for _, v := range nums {
		select {


		case out <- v * v:
		case <-done:
			return
		}
	}
	close(out)
}()
	return  out
}

func produce(done <-chan struct{},in <-chan int,str string) <-chan int{
	out := make(chan int)
	go func(){
		for val := range in {
			select {

				case out <- val:fmt.Println("Read by :" + str)
					case <-done:
						return
			}
		}
		close(out)
	}()
return out
}

func merge(done <-chan struct{},ch... <-chan int) <-chan int{
	var wg sync.WaitGroup
	out := make(chan int)
	output:= func(in <-chan int ){
		for v:= range in {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
		defer wg.Done()
	}

	wg.Add(len(ch))
	for _,v:=range ch{
		go output(v)
	}

	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}