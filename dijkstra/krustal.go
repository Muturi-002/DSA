/*
This is a modified version of graphs.go.
As the file's name suggests, this code seeks to create a weighted graph successfully.
*/
package main

import (
	"errors"
	"fmt"
	"strings"
	"sort"
)
type Graph struct {
	vertices []*Vertex
	edges []*Edge
}
type Vertex struct {
	key string
	neighbours []*Vertex//adjacent vertices conneected by an edge
	w []*Edge
}
type Edge struct {
	weight int//weight of an edge
}
func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %s not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}
func contains(s []*Vertex, k string) bool {
	for _,v:=range s {
		if k==v.key {
			return true
		}
	}
	return false
}
func (g *Graph) AddEdge(from, to string, distance int) { // the variables represent the root and destination vertices respectively
	from=strings.ToUpper(from)
	to=strings.ToUpper(to)
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	if fromVertex == nil || toVertex == nil { // check if any or all of the vertices do not exist
		err := errors.New("'invalid edge creation'")
		fmt.Print(err, from, "-->", to, ".")
		if fromVertex == nil {
			fmt.Print("Vertex ", from, " does not exist\n")
		} else {
			fmt.Print("Vertex ", to, " does not exist\n")
		}
		panic(err)
	} else if contains(fromVertex.neighbours, to) { // checks if there is an existing edge between the two vertices
		err := errors.New("'Edge already exists'")
		fmt.Print(err, from, "-->", to, "\n")
		panic(err)
	}

	fromVertex.neighbours = append(fromVertex.neighbours, toVertex)
	// Add the edge with weight
	edge := Edge{weight: distance}
	fromVertex.w = append(fromVertex.w, &edge)//dereference the Edge{} pointer
	g.edges = append(g.edges, &edge)//dereference the Edge{} pointer
	fmt.Printf("Edge added: %s --(%d)--> %s\n", from, edge.weight, to)
}
func (g *Graph) getVertex(k string) *Vertex { //checks for a vertex in the graph
	for i, v:=range g.vertices{
		if v.key==k{
			return g.vertices[i]
		}
	}
	return nil
}
func (g Graph) Print() {
	// print out each vertex and its neighbours with weights
	for _, v := range g.vertices {
		fmt.Print("Vertex ", v.key, ": ")
		for i, n := range v.neighbours {
			// Find the edge weight
			if i < len(v.w) {//the weight of the edge is represented by w, from the Edge struct
				fmt.Printf("(%d) %v ", v.w[i].weight, n.key)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func Krustal(graph [][]int, g *Graph) {
	fmt.Println("====Krustal's Algorithm====")
	// Initialize the set of edges and sort them by weight
	fmt.Println("Graph matrix:")
	for _,row:=range graph{
		fmt.Println(row)
	}
	type Cost struct {
		from, to, weight int
	}
	var edges []Cost
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] != 0 {
				edges = append(edges, Cost{from: i, to: j, weight: graph[i][j]})
			}
		}
	}
	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// Initialize union-find structure
	parent := make([]int, len(graph))
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(v int) int {
		if parent[v] != v {
			parent[v] = find(parent[v])
		}
		return parent[v]
	}

	union := func(v1, v2 int) {
		root1 := find(v1)
		root2 := find(v2)
		if root1 != root2 {
			parent[root2] = root1
		}
	}

	// Kruskal's algorithm
	var mst []Cost
	for _, edge := range edges {
		if find(edge.from) != find(edge.to) {
			mst = append(mst, edge)
			union(edge.from, edge.to)
		}
	}

	// Print the MST
	fmt.Println("Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("%d --(%d)--> %d\n", edge.from, edge.weight, edge.to)
	}
}

func main(){
	defer fmt.Println("---Weighted Graphs---")
	g := Graph{}
	vertices:=[]string{"A","B","C","D","E"}
	for _,r:=range vertices{
		g.AddVertex(r)
	}
	fmt.Println("Vertices of graph:")
	g.AddEdge("A", "B", 5)
	g.AddEdge("A", "C", 10)
	g.AddEdge("B", "C", 6)
	g.AddEdge("C", "D", 15)
	g.AddEdge("D", "E", 20)
	g.AddEdge("E", "A", 2)
	g.AddEdge("B", "D", 7)
	g.Print()
	graph:= [][]int{
		{0,5,10,0,0},
		{0,0,6,7,0},
		{0,0,0,15,0},
		{0,0,0,0,20},
		{2,0,0,0,0},
	}
	Krustal(graph,&g)
}