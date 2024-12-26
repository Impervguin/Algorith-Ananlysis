package graph

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type WeightedUndirectedGraph struct {
	nodesCount int
	edges      [][]float64
}

type WeightedCycle struct {
	nodes            []int
	weight           float64
	weightCalculated bool
	graph            *WeightedUndirectedGraph
}

// NewWeightedUndirectedGraph creates a new weighted undirected graph with the given number of nodes.
func NewWeightedUndirectedGraph(nodesCount int) *WeightedUndirectedGraph {
	graph := &WeightedUndirectedGraph{
		nodesCount: nodesCount,
		edges:      make([][]float64, nodesCount),
	}
	for i := range graph.edges {
		graph.edges[i] = make([]float64, nodesCount)
		for j := range graph.edges[i] {
			graph.edges[i][j] = math.Inf(1)
		}
	}
	return graph
}

// AddEdge adds a weighted edge between the given nodes with the specified weight.
func (g *WeightedUndirectedGraph) AddEdge(from, to int, weight float64) error {
	if from < 0 || from > g.nodesCount {
		return fmt.Errorf("invalid node")
	}
	if to < 0 || to > g.nodesCount {
		return fmt.Errorf("invalid node")
	}
	g.edges[from][to] = weight
	g.edges[to][from] = weight
	return nil
}

// GetEdgeWeight returns the weight of the edge between the given nodes.
func (g *WeightedUndirectedGraph) GetEdgeWeight(from, to int) (float64, error) {
	if from < 0 || from >= g.nodesCount {
		return 0, fmt.Errorf("invalid node")
	}
	if to < 0 || to >= g.nodesCount {
		return 0, fmt.Errorf("invalid node")
	}
	return g.edges[from][to], nil
}

func ReadWeightedUndirectedGraphCSV(filename string) (*WeightedUndirectedGraph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")
	gr := NewWeightedUndirectedGraph(len(parts))
	i := 0
	for j, part := range parts {
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}
		gr.AddEdge(i, j, num)
	}
	i++

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != gr.GetNodesCount() {
			return nil, fmt.Errorf("invalid CSV format")
		}
		for j, part := range parts {
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return nil, err
			}
			gr.AddEdge(i, j, num)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return gr, nil
}

func (g *WeightedUndirectedGraph) GetNodesCount() int {
	return g.nodesCount
}

func (g *WeightedUndirectedGraph) GetNodes() []int {
	nodes := make([]int, 0, g.nodesCount)
	for i := 0; i < g.nodesCount; i++ {
		nodes = append(nodes, i)
	}
	return nodes
}

func (g *WeightedUndirectedGraph) GetRandomHamiltonian() *WeightedCycle {
	nodes := g.GetNodes()
	rand.Shuffle(g.GetNodesCount(), func(i, j int) { nodes[i], nodes[j] = nodes[j], nodes[i] })
	cycle := NewWeightedCycle(g)

	for _, node := range nodes {
		err := cycle.AddNode(node)
		if err != nil {
			return nil
		}
	}
	return cycle
}

func NewWeightedCycle(graph *WeightedUndirectedGraph) *WeightedCycle {
	return &WeightedCycle{
		graph:            graph,
		weightCalculated: false,
	}
}

func (c *WeightedCycle) AddNode(node int) error {
	_, err := c.graph.GetEdgeWeight(node, node)
	if err != nil {
		return err
	}
	c.nodes = append(c.nodes, node)
	c.weightCalculated = false
	return nil
}

func (c *WeightedCycle) CalculateWeight() float64 {
	if c.weightCalculated {
		return c.weight
	}
	weight := 0.0
	for i := range c.nodes {
		w, err := c.graph.GetEdgeWeight(c.nodes[i], c.nodes[(i+1)%len(c.nodes)])
		if err != nil {
			return 0
		}
		weight += w
	}
	c.weight = weight
	c.weightCalculated = true
	return weight
}

func (c *WeightedCycle) GetNodes() []int {
	return c.nodes
}

func (c *WeightedCycle) LastNode() (int, error) {
	if len(c.nodes) == 0 {
		return -1, fmt.Errorf("empty cycle")
	}
	return c.nodes[len(c.nodes)-1], nil

}
