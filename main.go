package main

import (
	"graphModule/graphinit"
)

func main() {
	newGraph := new(graphinit.Graph)
	newGraph.AddNode("A", 1)
	newGraph.AddNode("B", 2)
	newGraph.AddNode("C", 3)
	newGraph.AddEdge("Edge1", "A", "B", 23)
	newGraph.AddEdge("Edge2", "B", "C", 46)
	newGraph.AddEdge("Edge3", "C", "A", 124)
	newGraph.WriteToJson("graph.json")
}
