package main

import (
	"fmt"
	"sort"
)
const Infinity = int(^uint(0) >> 1)
type Node struct {
	val string
}
type Edge struct {
	Parent *Node
	child *Node
	weight int
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddEdge(parent,child *Node,cost int){
	edge := &Edge{parent,child,cost}
	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

func (g *Graph) AddNode(n *Node){
	var isPresent bool
	for _,v := range g.Nodes{
		if  v==n{
			isPresent =true
		}
	}

	if !isPresent{
		g.Nodes = append(g.Nodes, n)
	}
}

func (g *Graph) GetNodeEdges(start *Node)(edges []*Edge){

	for _,edge:= range g.Edges{
		if edge.Parent==start {
			edges = append(edges, edge)
		}
	}
	return edges
}



func (g *Graph) Djikstra(startNode *Node) (shortestPathTable string){
	// intialiaze all node to zero except the first node
	costTable := g.InitialiazeAllNodes(startNode)

	var visitedNode []*Node
	for len(g.Nodes)!= len(visitedNode){
		// Get closest non visited Node (lower cost) from the costTable
		node := getClosestNonVisitedNode(costTable, visitedNode)
		// Mark Node as visited
		visitedNode = append(visitedNode, node)
		nodeEdges := g.GetNodeEdges(node)
		for _,edge := range nodeEdges{
			distanceToNeighbor := costTable[node]+edge.weight
			if distanceToNeighbor<costTable[edge.child]{
				costTable[edge.child] = distanceToNeighbor
			}
		}
	}

	// Make the costTable nice to read :)
	for node, cost := range costTable {
		shortestPathTable += fmt.Sprintf("Distance from %s to %s = %d\n", startNode.val, node.val, cost)
	}

	return shortestPathTable
}
// getClosestNonVisitedNode returns the closest Node (with the lower cost) from the costTable
// **if the node hasn't been visited yet**
func getClosestNonVisitedNode(costTable map[*Node]int, visited []*Node) *Node {
	type CostTableToSort struct {
		Node *Node
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, CostTableToSort{node, cost})
		}
	}

	// We need the Node with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}
func (g *Graph) InitialiazeAllNodes(startNode *Node) map[*Node]int{
	constable := make(map[*Node]int)
	constable[startNode] = 0
	for _,node := range g.Nodes{
		if node != startNode {
			constable[node] = Infinity
		}
	}
	return constable
}

func main(){
	a:= &Node{"A"}
	b:= &Node{"B"}
	c:= &Node{"C"}
	d:= &Node{"D"}
	e:=  &Node{"E"}
	f := &Node{"F"}
	g := &Node{"G"}

	graph := Graph{}
	graph.AddEdge(a, c, 2)
	graph.AddEdge(a, b, 5)
	graph.AddEdge(c, b, 1)
	graph.AddEdge(c, d, 9)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(d, e, 2)
	graph.AddEdge(d, g, 30)
	graph.AddEdge(d, f, 10)
	graph.AddEdge(f, g, 1)
	fmt.Println(graph.Djikstra(a))
}