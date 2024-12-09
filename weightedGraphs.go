/*This is a modified version of graphs.go.
As the file's name suggests, this code seeks to create a weighted graph successfully.
*/
package main
import (
	"fmt"
	"errors"
	"strconv"
)
type Graph struct {
	vertices []*Vertex
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
		vKey:=strconv.Itoa(v.key)
		if k==vKey {
			return true
		}
	}
	return false
}
func (g *Graph) AddEdge(from, to string, distance int) { // the variables represent the root and destination vertices respectively
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
	fmt.Printf("Edge added: %s --(%d)--> %s\n", from, edge.weight, to)
}
func (g *Graph) getVertex(k string) *Vertex { //checks for a vertex in the graph
	for i, v:=range g.vertices{
		vKey:=strconv.Itoa(v.key)
		if vKey==k{
			return g.vertices[i]
		}
	}
	return nil
}
func (g *Graph) Print() {
	// print out each vertex and its neighbours with weights
	for _, v := range g.vertices {
		fmt.Print("Vertex ", v.key, ": ")
		for _, n := range v.neighbours {
			// Find the edge weight
			for _, edge := range g.edges {
				if (edge.from == v.key && edge.to == n.key) || (edge.from == n.key && edge.to == v.key) {
					fmt.Printf("%v (%d) ", n.key, edge.weight)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}


func main(){
	defer fmt.Println("---Weighted Graphs---")
	g := Graph{}
	vertices:=[]string{"A","B","C","D","E"}
	for _,r:=range vertices{
		g.AddVertex(r)
	}
	fmt.Println("Vertices of graph:")
	fmt.Println(g)
	g.Print()
	g.AddEdge("A", "B", 5)
	g.AddEdge("A", "C", 10)
	g.AddEdge("B", "C", 6)
	g.AddEdge("C", "D", 15)
	g.AddEdge("C", "D", 20)
	g.Print()
}