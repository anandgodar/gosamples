package main

import "fmt"

func main(){
	arr := []int{18,12,4,10,30,50,2,7,89,98,77}
	fmt.Println(QuickSort(arr,0,len(arr)-1))
}
func QuickSort(arr []int,l,h int)[]int{
	if l<h{
		j := partition(arr,l,h)
		QuickSort(arr,l,j-1)
		QuickSort(arr,j+1,h)
	}
	return arr
}
func partition(arr []int,l,h int)int{

	pivot := arr[h]
	i:=l-1

	for j:=l;j<h;j++{
		if arr[j]<pivot{
			i++
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
	arr[i+1],arr[h]=arr[h],arr[i+1]
	return i+1
}
