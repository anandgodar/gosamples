package main

import "fmt"

type Tree struct {
	Value int
	Repeat int
	Left *Tree
	Right *Tree
	Parent *Tree
}

func (t *Tree) Insert(v int) error {

	switch {
	case t.Value == v:
		t.Repeat ++
	case t.Value<v :
		if t.Left==nil{
			t.Left = &Tree{v,0,nil,nil,nil}
			return nil
		}
		return t.Left.Insert(v)
	case t.Value>v:
		if t.Right==nil{
			t.Right = &Tree{v,0,nil,nil,nil}
		}
		return t.Right.Insert(v)
	}

	return nil
}
func BinaryTree(nums []int) *Tree{
	vroot:= nums[0]
	tree := Tree{vroot,0,nil,nil,nil}
	for _,value := range nums[1:]{
		tree.Insert(value)
	}
	return &tree
}

func BFS(tree *Tree)[]int{
	queue := []*Tree{}
	result:= []int{}
	queue = append(queue,tree)
	return BFSUtil(queue,result)
}
func BFSUtil(queue []*Tree,res []int)[]int{
	if len(queue)==0{
		return res
	}
	res = append(res,queue[0].Value)
	if queue[0].Right !=nil{
		queue = append(queue,queue[0].Right)
	}
	if queue[0].Left !=nil{
		queue = append(queue,queue[0].Left)
	}
	return BFSUtil(queue[1:],res)
}

func (t *Tree) Traverse(n *Tree, f func(*Tree)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)

}
func main(){
	nums := []int{10,2,5,4,7,8,1,34,22}
	tree:=BinaryTree(nums)
	//fmt.Println(BFS(tree))

	tree.Traverse(tree, func(n *Tree) { fmt.Print(n.Value,",") })

}