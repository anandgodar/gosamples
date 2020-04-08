package main

import (
	"fmt"
	"testing"
)

func sum(s []int) int {

	if len(s)==0{
		return 0
	}
	x, s := s[len(s)-1], s[:len(s)-1]
	fmt.Println(x, s)
	return x + sum(s)
}
func main(){
	arr := []int{4,5,6,7}
	fmt.Println(sum(arr))
}
func Test_Sum(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "sum #1",
			input:  []int{3, 4, 5},
			output: 12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := sum(tC.input)

			if r != tC.output {
				t.Errorf("expected %d, got %d", tC.output, r)
			}
		})
	}
}