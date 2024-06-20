package graphinit

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Node struct {
	Name string `json:"Name"`
	ASN  int    `json:"Asn"`
}

type Edge struct {
	NodeA     string `json:"NodeA"`
	NodeB     string `json:"NodeB"`
	IgpMetric int    `json:"IgpMetric"`
}

type Graph struct {
	Nodes map[string]Node
	Edges map[string]Edge
}

func LoadFromJson(filename string) Graph {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error occurred when opening file %s : %s", filename, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error occured when pulling bytes from file %s: Error: %s", filename, err)
	}

	var graph Graph
	if err := json.Unmarshal(bytes, &graph); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	fmt.Println(graph)
	return graph
}

func (g *Graph) AddNode(name string, asn int) {
	if g.Nodes == nil {
		g.Nodes = make(map[string]Node)
	}
	if _, ok := g.Nodes[name]; !ok {
		newNode := Node{
			Name: name,
			ASN:  asn,
		}
		g.Nodes[name] = newNode
	} else {
		fmt.Printf("Node with name %s already exists! ", name)
	}

}

func (g *Graph) findNode(name string) bool {
	_, ok := g.Nodes[name]
	return ok
}

func (g *Graph) AddEdge(name string, nodeA string, nodeB string, igpMetric int) error {
	if g.Edges == nil {
		g.Edges = make(map[string]Edge)
	}
	if _, ok := g.Edges[name]; !ok {
		if ok := g.findNode(nodeA); !ok {
			return fmt.Errorf("Node %s does not exist in the graph", nodeA)
		}
		if ok := g.findNode(nodeB); !ok {
			return fmt.Errorf("Node %s does not exist in the graph", nodeB)
		}
		newEdge := Edge{
			NodeA:     nodeA,
			NodeB:     nodeB,
			IgpMetric: igpMetric,
		}
		g.Edges[name] = newEdge
	} else {
		fmt.Printf("Edge %s already exists in the graph!", name)
	}

	return nil
}

func (g *Graph) WriteToJson(filename string) {
	updatedBytes, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %s", err)
	}

	// Write the updated JSON back to the file.
	if err := os.WriteFile(filename, updatedBytes, 0644); err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}

	fmt.Println("Graph updated successfully.")
}
