// Description: This file contains the implementation of Dijkstra's algorithm for finding the shortest path in a graph.
package dijkstra
import (
	"fmt"
	"github.com/Muturi-002/DSA../dijkstra/graphweights"
)
var visited bool
func Dijkstra(g *graphweights.Graph,startkey string) {
	const infinity int= 9999
	fmt.Println("====Dijkstra's Algorithm====")
	start := g.GetVertex(startkey)
	if start==nil{
		fmt.Println("Vertex not found")
		return
	}
	//initialize the distance of the start vertex to 0
	start.w[0].weight=0
	//initialize the distance of all other vertices to infinity
	for _,v:=range g.vertices{
		if v.key!=startkey{
			v.w[0].weight=infinity
		}
	}
	//initialize the set of visited vertices to empty
	visited=false
	//initialize the set of unvisited vertices to all vertices
	unvisited:= g.vertices
	for len(unvisited)>0{
		//select the unvisited vertex with the smallest distance
		u:=unvisited[0]
		for _,v:=range unvisited{
			if v.w[0].weight<u.w[0].weight{
				u=v
			}
		}
		//remove the selected vertex from the set of unvisited vertices
		unvisited=remove(unvisited,u)
		//for each of the vertex's unvisited neighbours, calculate their tentative distances through the current vertex
		for i:=0;i<len(u.neighbours);i++{
			v:=u.neighbours[i]
			if graphweights.Contains(unvisited,v.key){
				alt:=u.w[0].weight+u.w[i].weight
				if alt<v.w[0].weight{
					v.w[0].weight=alt
				}
			}
		}
	}
	//when we are done, print the distances
	for _,v:=range g.vertices{
		fmt.Printf("Vertex %s: %d\n",v.key,v.w[0].weight)
	}
}
func remove(s []*graphweights.Vertex, v *graphweights.Vertex) []*graphweights.Vertex {
	var r []*graphweights.Vertex
	for _,v:=range s{
		if v!=v{
			r=append(r,v)
		}
	}
	return r
}