package main

import (
	"fmt"
	"lab6/internal/ant"
	"lab6/internal/graph"
	"os"
)

type Benchmark struct {
	MinPath float64
	MaxPath float64
	AvgPath float64
}

const oneGraphCounts = 10

const antCount = 10
const eliteAntCount = 5
const initPheromone = 0.01
const pheromonePerAnt = 0.1

func main() {
	distCoeffs := []float64{0.1, 0.3, 0.5, 0.7, 0.9}
	evapCoeffs := []float64{0.1, 0.3, 0.5, 0.7, 0.9}
	daysCounts := []int{5, 20, 50, 100, 200}
	f, err := os.Create("reslatex.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gr1, err := graph.ReadWeightedUndirectedGraphCSV("gr1.csv")
	if err != nil {
		panic(err)
	}
	gr2, err := graph.ReadWeightedUndirectedGraphCSV("gr2.csv")
	if err != nil {
		panic(err)
	}
	gr3, err := graph.ReadWeightedUndirectedGraphCSV("gr3.csv")
	if err != nil {
		panic(err)
	}
	grs := []*graph.WeightedUndirectedGraph{gr1, gr2, gr3}
	for _, dist := range distCoeffs {
		for _, evap := range evapCoeffs {
			for _, days := range daysCounts {
				b := make([]Benchmark, 0, len(grs))
				for _, gr := range grs {
					algorithm := ant.NewElitistAntAlgorithm(dist, 1-evap, evap, initPheromone, pheromonePerAnt, antCount, eliteAntCount, days)
					bestCycle, err := algorithm.Run(gr)
					if err != nil {
						panic(err)
					}
					minPath := bestCycle.CalculateWeight()
					maxPath := minPath
					sumPath := minPath
					for r := 1; r < oneGraphCounts; r++ {
						bestCycle, err := algorithm.Run(gr)
						if err != nil {
							panic(err)
						}
						path := bestCycle.CalculateWeight()
						minPath = min(minPath, path)
						maxPath = max(maxPath, path)
						sumPath += path
					}
					avgPath := sumPath / oneGraphCounts
					b = append(b, Benchmark{minPath, maxPath, avgPath})
				}
				fmt.Fprintf(f, "%.1f & %.1f & %d", dist, evap, days)
				for _, bench := range b {
					fmt.Fprintf(f, " & %.2f & %.2f & %.2f", bench.MinPath/100, bench.MaxPath/100, bench.AvgPath/100)
				}
				fmt.Fprintln(f, "\\\\")
				fmt.Fprintln(f, "\\hline")
			}
		}
	}

}
