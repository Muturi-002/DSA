/*This is a copy of weightedGraphs.go in the root directory. It is specific to the dijkstra directory, meant to be imported to the dijkstra.go file.
 */
package main

import (
	"errors"
	"fmt"
	//d "github.com/Muturi-002/DSA/dijkstra/d2"
)
const infinity int= 99999
var vertCount int//number of vertices in the graph
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
func (g *Graph) addVertex(k string) {
	if Contains(g.vertices, k) {
		err := fmt.Errorf("vertex %s not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}
func Contains(s []*Vertex, k string) bool {
	for _,v:=range s {
		if k==v.key {
			return true
		}
	}
	return false
}
func (g *Graph) addEdge(from, to string, distance int) { // the variables represent the root and destination vertices respectively
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
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
	} else if Contains(fromVertex.neighbours, to) { // checks if there is an existing edge between the two vertices
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
func (g *Graph) GetVertex(k string) *Vertex { //checks for a vertex in the graph
	for i, v:=range g.vertices{
		if v.key==k{
			return g.vertices[i]
		}
	}
	return nil
}
func (g Graph) print() {
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
func Dijkstra(g *Graph,startkey string) {
	fmt.Println("====Dijkstra's Algorithm====")
	distance := make([]int, len(g.vertices))//creates an array for the shortest distance from the source vertex to the neighbouring vertices
    visited := make([]bool, len(g.vertices))
	vertCount= len(g.vertices)
	for k := 0; k < vertCount; k++ {
        distance[k] = infinity
        visited[k] = false
    }

	start := g.GetVertex(startkey)
	if start == nil {
		fmt.Println("Vertex not found")
		return
	}
	var startIndex int//enables the distance between the start vertex and itself to be zero. Setting the distance index as start will create an error.
	for i, v := range g.vertices {
		if v == start {
			startIndex = i
			break
		}
	}
	distance[startIndex] = 0

    for k := 0; k < vertCount; k++ {
        m := miniDist(distance, visited)
        visited[m] = true

        for l := 0; l < vertCount; l++ {
			for _, edge := range g.vertices[m].w {
				//neighbour := g.vertices[m].neighbours[k]
				if !visited[l] && edge.weight != 0 && distance[m] != infinity &&
					distance[m]+edge.weight < distance[l] {
					distance[k] = distance[m] + edge.weight
				}
			}
            }
        }
    fmt.Println("Vertex\t\tDistance from source vertex")
    for k := 0; k < len(g.vertices); k++ {
        fmt.Printf("%s\t\t\t%d\n", g.vertices[k], distance[k])
    }
}
func miniDist(distance []int, visited []bool) int {// a method that seeks to find the smallest distance from the source vertex to the neighbouring vertices
    minimum := infinity
    var ind int// stores the index of the vertex with the smallest distance
    for k := 0; k < 6; k++ {
        if !visited[k] && distance[k] < minimum {
            minimum = distance[k]
            ind = k
        }
    }
    return ind
}
func main(){
	fmt.Println("---Weighted Graphs---")
	g := Graph{}
	vertices:=[]string{"A","B","C","D","E"}
	for _,r:=range vertices{
		g.addVertex(r)
	}
	fmt.Println("Vertices of graph:")
	g.addEdge("A", "B", 5)
	g.addEdge("A", "C", 10)
	g.addEdge("B", "C", 6)
	g.addEdge("C", "D", 15)
	g.addEdge("D", "E", 20)
	g.print()
	var src string
	fmt.Printf("Choose the start vertex: ")
	fmt.Scanf("%s", &src)
	fmt.Println("\nShortest path from ", src, " to other vertices:")
	Dijkstra(&g, src)
}
