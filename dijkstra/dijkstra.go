package dijkstra

import (
	"fmt"
	"./weightedGraphs.go"
)
const visited bool
src string
func (g *Graph) Dijkstra() {
	fmt.Println("====Dijkstra's Algorithm====")
	fmt.Printf("Choose the start vertex: ")
	fmt.Scanf(&src)
	fmt.Println()
	fmt.Println("The shortest path from the source vertex",src," to the other vertices is as follows:")
	
}