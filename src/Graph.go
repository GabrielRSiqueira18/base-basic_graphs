package src

import "fmt"

type Edge struct {
	Src, Dest int
}

type Graph struct {
	V    int
	E    int
	Edge []Edge
}

func CreateGraph(V, E int) *Graph {
	return &Graph{Edge: make([]Edge, E), V: V, E: E}
}

func (graph *Graph) Print() {
	edges := graph.Edge
	for i := 0; i < len(edges); i++ {
		fmt.Printf("Edge %d: %d->%d\n", i+1, edges[i].Src, edges[i].Dest)
	}
}

func (graph *Graph) CreateEdges() {
	var src, dest int
	exists := false
	for i := 0; i <= graph.V; i++ {
		fmt.Printf("Digite o vértice do começo da aresta: %v: ", i+1)
		fmt.Scanf("%d\n", &src)
		fmt.Printf("Digite o vértice do final da aresta: %v: ", i+1)
		fmt.Scanf("%d\n", &dest)

		for _, g := range graph.Edge {
			if g.Dest == dest && g.Src == src {
				fmt.Println("Aresta já existe")

				exists = true

				for exists {
					fmt.Printf("Digite o vértice do começo da aresta: %v: ", i+1)
					fmt.Scanf("%d\n", &src)
					fmt.Printf("Digite o vértice do final da aresta: %v: ", i+1)
					fmt.Scanf("%d\n", &dest)

					if g.Dest != dest && g.Src != src {
						exists = false
						break
					}

					fmt.Println("Aresta já existe")
				}

			}
		}

		graph.Edge[i].Src = src
		graph.Edge[i].Dest = dest
	}
}
