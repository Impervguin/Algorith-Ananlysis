package ant

import "lab6/internal/graph"

type GraphWithPheromon struct {
	graph     *graph.WeightedUndirectedGraph
	pheromone [][]float64
}

func NewGraphWithPheromon(graph *graph.WeightedUndirectedGraph, pheromoneInit float64) *GraphWithPheromon {
	graphph := &GraphWithPheromon{
		graph:     graph,
		pheromone: make([][]float64, graph.GetNodesCount()),
	}
	for i := range graphph.pheromone {
		graphph.pheromone[i] = make([]float64, graph.GetNodesCount())
		for j := range graphph.pheromone[i] {
			graphph.pheromone[i][j] = pheromoneInit
		}
	}
	return graphph
}

func (graph *GraphWithPheromon) ApplyPheromon(path *graph.WeightedCycle, pheromonPerAnt float64) {
	length := path.CalculateWeight()
	pheromonPerlength := pheromonPerAnt / length / float64(len(path.GetNodes()))
	nodes := path.GetNodes()
	for i := range nodes {
		graph.pheromone[nodes[i]][nodes[(i+1)%len(nodes)]] += pheromonPerlength
		graph.pheromone[nodes[(i+1)%len(nodes)]][nodes[i]] += pheromonPerlength
	}
}

func (graph *GraphWithPheromon) EvaporatePheromone(phcoeff float64) {
	for i := range graph.pheromone {
		for j := range graph.pheromone[i] {
			graph.pheromone[i][j] *= (1 - phcoeff)
		}
	}
}
