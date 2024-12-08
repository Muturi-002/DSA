package main
import (
	"fmt"
	"errors"
)
type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key int
	neighbours []*Vertex//adjacent vertices conneected by an edge
	weight int//shows weight of edge
}
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}
func (g *Graph) AddEdge(from, to, distance int){//the variables represent the root and destination vertices respectively
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()
	if fromVertex == nil || toVertex == nil {//check if any or all of the vertices vertices do not exist
		err := errors.New("'invalid edge creation'")
		fmt.Print(err,from,"-->",to,".")
		if fromVertex == nil {
			fmt.Print("Vertex ",from," does not exist\n")
		}else{
			fmt.Print("Vertex ",to," does not exist\n")
		}
		panic(err)	
	}else if contains(fromVertex.neighbours, to) {//checks if there is an existing edge between the two vertices
		err := errors.New("'Edge already exists'")
		fmt.Print(err,from,"-->",to,"\n")
		panic(err)	
	}
	
	fromVertex.neighbours = append(fromVertex.neighbours, toVertex)
	
}
func (g *Graph) Print(){
	//print out each vertex and its neighbours
	for _,v:=range g.vertices {
		fmt.Print("Vertex ",v.key,": ")
		for _,v:=range v.neighbours {
			fmt.Printf("%v (%v)",v.key,v.neighbours.weight)
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Graph) getVertex(k int) *Vertex { //checks for a vertex in the graph
	for i, v:=range g.vertices{
		if v.key==k{
			return g.vertices[i]
		}
	}
	return nil
}
func contains(s []*Vertex, k int) bool {
	for _,v:=range s {
		if k==v.key {
			return true
		}
	}
	return false
}
func main(){
	defer fmt.Println("---Weighted Graphs---")
	g := Graph{}
	for i:=0;i<5;i++ {
		g.AddVertex(i)
	}
	fmt.Println("Vertices of graph:")
	fmt.Println(g)
	g.Print()
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 10)
	g.AddEdge(1, 2, 6)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 20)
	g.Print()
}