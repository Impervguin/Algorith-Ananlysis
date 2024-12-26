package ant

import (
	"fmt"
	"lab6/internal/graph"
	"math"
	"math/rand"
)

type Ant struct {
	pheromone      float64
	distanceCoeff  float64
	pheromoneCoeff float64

	path      *graph.WeightedCycle
	graph     *GraphWithPheromon
	unvisited []int
}

// NewAnt creates a new ant with the given graph and initial pheromone level.
func NewAnt(gr *GraphWithPheromon, initialPheromone, distanceCoeff, pheromoneCoeff float64, startNode int) *Ant {
	a := &Ant{
		pheromone:      initialPheromone,
		distanceCoeff:  distanceCoeff,
		pheromoneCoeff: pheromoneCoeff,
		graph:          gr,
		unvisited:      gr.graph.GetNodes(),
	}
	a.path = graph.NewWeightedCycle(gr.graph)
	a.path.AddNode(startNode)
	a.unvisited = append(a.unvisited[:startNode], a.unvisited[startNode+1:]...)
	return a
}

func (a *Ant) desireFunc(node1, node2 int) float64 {
	edgeWeight, _ := a.graph.graph.GetEdgeWeight(node1, node2)
	distanceDesire := 1 / edgeWeight
	pheromone := a.graph.pheromone[node1][node2]

	return math.Pow(distanceDesire, a.distanceCoeff) * math.Pow(pheromone, a.pheromoneCoeff)
}

func (a *Ant) chooseNextNode() int {
	lastNode, _ := a.path.LastNode()

	sumDesire := 0.
	for _, node := range a.unvisited {
		sumDesire += a.desireFunc(lastNode, node)
	}

	probabilities := make(map[int]float64, len(a.unvisited))
	for i := range a.unvisited {
		probabilities[i] = a.desireFunc(lastNode, a.unvisited[i]) / sumDesire
	}
	randVal := rand.Float64()
	sumProb := 0.
	for i, prob := range probabilities {
		sumProb += prob
		if sumProb >= randVal {
			node := a.unvisited[i]
			a.unvisited = append(a.unvisited[:i], a.unvisited[i+1:]...)
			return node
		}
	}
	return -1
}

func (a *Ant) Go() error {
	for len(a.unvisited) > 0 {
		nextNode := a.chooseNextNode()
		if nextNode == -1 {
			return fmt.Errorf("no valid next node found")
		}
		a.path.AddNode(nextNode)
	}
	return nil
}

func (a *Ant) GetPath() *graph.WeightedCycle {
	return a.path
}

func (a *Ant) GetPheromone() float64 {
	return a.pheromone
}
