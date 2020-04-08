package main

import "fmt"

func main(){
	var sm int
	nums := []int{13,19,3,45,6,7}
	for i:=0;i<len(nums)-1;i++{
			sml := nums[i]
			index := i
			for j:=i+1;j<len(nums);j++{
				if sml>nums[j]{
					sm = j
					sml = nums[j]
				}
			}
			if sml != nums[i] {
				nums[index], nums[sm] = nums[sm], nums[index]
			}
	}

	fmt.Println(nums)
}
