//Breadth First Search algorithm
package main
import (
	"fmt"
	"errors"
	"strings"
)
//Creation of graph
type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key string
	neighbours []*Vertex
}
//functions related to the Graph struct
func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}
func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	if fromVertex == nil || toVertex == nil {
		err := errors.New("invalid edge creation")
		fmt.Print(err, from, "-->", to, ".")
		if fromVertex == nil {
			fmt.Print("Vertex ", from, " does not exist\n")
		} else {
			fmt.Print("Vertex ", to, " does not exist\n")
		}
		panic(err)
	} else if contains(fromVertex.neighbours, to) {
		err := errors.New("Edge already exists")
		fmt.Print(err, from, "-->", to, "\n")
		panic(err)
	}
	fromVertex.neighbours = append(fromVertex.neighbours, toVertex)
}

func (g *Graph) display() {
	for _, v := range g.vertices {
		fmt.Print("Vertex ", v.key, ": ")
		for _, v := range v.neighbours {
			fmt.Print(v.key, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Graph) getVertex(k string) *Vertex {
	for i, v:=range g.vertices{
		if v.key==k{
			return g.vertices[i]
		}
	}
	return nil
}
func contains(s []*Vertex, k string) bool {
	for _,v:=range s {
		if k==v.key {
			return true
		}
	}
	return false
}
//Breadth First Search algorithm
func (g *Graph) bfs(start string){
	q:=[]string{}
	visited:= make(map[string]bool)
	present:=g.getVertex(start)
	if present!=nil{
		fmt.Println("Starting from vertex: ",start)
		q= append(q,start)
		visited[start]=true
		for len(q)>0{
			current:=q[0]
			q= q[1:]//leaves out the current visited vertex(vertices) from being assessed
			fmt.Println(current)
			for _,v:=range g.getVertex(current).neighbours{
				if !visited[v.key]{
					q= append(q,v.key)
					visited[v.key]=true
				}
			}
		}
	}else{
		fmt.Println("Vertex does not exist")
	}
	fmt.Println()
}
func main(){
	var newGraph Graph
	vertices:= []string{"A","B","C","D","E","F","G","H","I","J"}
	for _,v:=range vertices{
		newGraph.AddVertex(v)
	}
	fmt.Println("Undirected Graph vertices: ")
	newGraph.AddEdge("A","B")
	newGraph.AddEdge("A","C")
	newGraph.AddEdge("B","D")
	newGraph.AddEdge("F","E")
	newGraph.AddEdge("B","F")
	newGraph.AddEdge("C","G")
	newGraph.AddEdge("C","H")
	newGraph.AddEdge("D","E")
	newGraph.AddEdge("D","J")
	newGraph.AddEdge("E","I")
	newGraph.AddEdge("J","H")
	newGraph.AddEdge("G","I")
	newGraph.AddEdge("H","J")
	newGraph.AddEdge("G","F")
	newGraph.AddEdge("H","A")
	newGraph.AddEdge("I","B")
	newGraph.display()
	fmt.Println("---Breadth First Search---")
	var start string
	fmt.Print("Enter the starting vertex: ")
	fmt.Scan(&start)
	newGraph.bfs(strings.ToUpper(start))
	
}