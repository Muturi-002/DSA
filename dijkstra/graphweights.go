/*This is a copy of weightedGraphs.go in the root directory. It is specific to the dijkstra directory, meant to be imported to the dijkstra.go file.
 */
package main

import (
	"errors"
	"fmt"
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
func Dijkstra(graph [5][5]int, g *Graph) {
	fmt.Println("====Dijkstra's Algorithm====")
	distance:= make([][]int, 5)//stores the distance from the source vertex to the other vertices
	visited := make([]bool, 5) // stores the visited vertices
	var src string
	var srcindex int
	fmt.Print("Enter the source vertex: ")
	fmt.Scan(&src); fmt.Println()
	for i,n:=range g.vertices{
		if n.key==src{
			srcindex=i
		}
	}
	fmt.Printf("Source vertex: %s. Index: %d\n", src, srcindex)
	for i := 0; i < 5; i++ {
		distance[i][i] = infinity
		visited[i] = false
	}

	distance[srcindex][srcindex] = 0

	for i := 0; i < 5-1; i++ {
		u := miniDist(distance, visited)
		visited[u] = true

		for v := 0; v < 5; v++ {
			if !visited[v] && graph[u][v] != 0 && distance[u][u] != infinity && distance[u][u]+graph[u][v] < distance[v][v] {
				distance[v][v] = distance[u][u] + graph[u][v]
			}
		}
	}

	fmt.Println("Vertex \t Distance from Source")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s \t %d\n", g.vertices[i].key, distance[i])
	}
}
func miniDist(distance [][]int, visited []bool) int {// a method that seeks to find the smallest distance from the source vertex to the neighbouring vertices
    minimum := infinity
    var ind int// stores the index of the vertex with the smallest distance
    for k := 0; k < 5; k++ {
		for j:=0;j<5;j++{
			if !visited[j] && distance[k][j] < minimum {
				minimum = distance[k][j]
				ind = k
			}
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
	g.addEdge("E", "B", 2)
	g.print()
	graph := [5][5]int{//from the graph defined
		{0, 5, 10, 0, 0},
		{5, 0, 6, 0, 0},
		{10, 6, 0, 15, 0},
		{0, 0, 15, 0, 20},
		{0, 2, 0, 20, 0},
	}
	fmt.Println("Graph matrix: ",graph)
	Dijkstra(graph,&g)
}
