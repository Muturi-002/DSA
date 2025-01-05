package main

import (
	"fmt"
	"sort"
)

func Kruskal(graph [][]int) {
	fmt.Println("====Kruskal's Algorithm====")
	// Initialize the set of edges and sort them by weight
	fmt.Println("Graph matrix:")
	for _, row := range graph {
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

func main() {
	defer fmt.Println("---Weighted Graphs---")
	graph := [][]int{
		{0, 5, 10, 0, 0},
		{0, 0, 6, 7, 0},
		{0, 0, 0, 15, 0},
		{0, 0, 0, 0, 20},
		{2, 0, 0, 0, 0},
	}
	Kruskal(graph)
}
