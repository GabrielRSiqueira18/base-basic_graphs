package main

import "app/src"

func main() {
	graph := src.CreateGraph(4, 5)

	graph.CreateEdges()
	graph.Print()
}
