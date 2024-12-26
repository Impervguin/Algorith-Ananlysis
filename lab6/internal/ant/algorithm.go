package ant

import (
	"lab6/internal/graph"
)

type ElitistAntAlgorithm struct {
	distanceCoeff    float64
	pheromoneCoeff   float64
	evaporationCoeff float64
	initPheromone    float64
	pheromonePerAnt  float64

	antsCount      int
	eliteAntsCount int
	daysCount      int
}

func NewElitistAntAlgorithm(distanceCoeff, pheromoneCoeff, evaporationCoeff, initPheromone, pheromonePerAnt float64, antsCount, eliteAntsCount, daysCount int) *ElitistAntAlgorithm {
	return &ElitistAntAlgorithm{
		distanceCoeff:    distanceCoeff,
		pheromoneCoeff:   pheromoneCoeff,
		evaporationCoeff: evaporationCoeff,
		initPheromone:    initPheromone,
		pheromonePerAnt:  pheromonePerAnt,

		antsCount:      antsCount,
		eliteAntsCount: eliteAntsCount,
		daysCount:      daysCount,
	}
}

func (a *ElitistAntAlgorithm) Run(gr *graph.WeightedUndirectedGraph) (*graph.WeightedCycle, error) {
	bestCycle := gr.GetRandomHamiltonian()

	phgr := NewGraphWithPheromon(gr, a.initPheromone)
	for day := 0; day < a.daysCount; day++ {
		ants := make([]*Ant, a.antsCount)
		initNode := gr.GetNodes()[0]
		for i := 0; i < a.antsCount; i++ {
			ants[i] = NewAnt(phgr, a.pheromonePerAnt, a.distanceCoeff, a.pheromoneCoeff, initNode)
		}

		for _, ant := range ants {
			if err := ant.Go(); err != nil {
				return nil, err
			}

			if ant.GetPath().CalculateWeight() < bestCycle.CalculateWeight() {
				bestCycle = ant.GetPath()
			}
		}

		phgr.EvaporatePheromone(a.evaporationCoeff)

		for _, ant := range ants {
			phgr.ApplyPheromon(ant.GetPath(), ant.pheromone)
		}

		for i := 0; i < a.eliteAntsCount; i++ {
			phgr.ApplyPheromon(bestCycle, a.pheromonePerAnt)
		}
	}
	return bestCycle, nil
}
