package main

import (
	"fmt"
	"lab6/internal/ant"
	"lab6/internal/fullsearch"
	"lab6/internal/graph"
)

func main() {
	gr, err := graph.ReadWeightedUndirectedGraphCSV("test2.csv")
	if err != nil {
		panic(err)
	}

	antColony := ant.NewElitistAntAlgorithm(
		0.5,
		0.9,
		0.2,
		0.01,
		0.1,
		5,
		2,
		15,
	)
	bestSolution, err := antColony.Run(gr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Best Ant solution: %v\n", bestSolution)
	fmt.Printf("Total Ant distance: %f\n", bestSolution.CalculateWeight())

	full := fullsearch.NewFullSearch()
	bestSolution, err = full.Run(gr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Best Full Search solution: %v\n", bestSolution)
	fmt.Printf("Total Full Search distance: %f\n", bestSolution.CalculateWeight())
}
