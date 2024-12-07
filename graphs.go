//from the youtube link https://youtu.be/D3oPCqn_HO8
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
}
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}
func (g *Graph) AddEdge(from, to int){//the variables represent the root and destination vertices respectively
	fromVertex:= g.getVertex(from)
	toVertex:=g.getVertex(to)
	err:= errors.New("Invalid edge creation of")
	//checking for errors
	//empty vertices
	defer recover()
	if fromVertex==nil || toVertex==nil{
		fmt.Println(err, "",from,"->",to)
		panic(err)
	}
	fromVertex.neighbours= append(fromVertex.neighbours,toVertex)

}
func (g *Graph) Print(){
	//print out each vertex and its neighbours
	for _,v:=range g.vertices {
		fmt.Print("Vertex ",v.key,": ")
		for _,v:=range v.neighbours {
			fmt.Print(v.key)
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Graph) getVertex(k int) *Vertex {
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
func main() {
	defer fmt.Println("---Graphs---")
	fmt.Println("creating a graph by using adjacency list")
	graph1:= &Graph{}
	for i:=0;i<5;i++ {
		graph1.AddVertex(i)
	}
	fmt.Println("Vertices in graph1 : ",*graph1)
	graph1.Print()
	graph1.AddVertex(5)
	graph1.AddEdge(1,2)
	graph1.AddEdge(7,5)
	graph1.Print()
}