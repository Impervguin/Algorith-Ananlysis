package fullsearch

import (
	"lab6/internal/graph"
)

func HeapAlgo(arr []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		c := make([]int, len(arr))
		for i := range c {
			c[i] = 0
		}
		arrCopy := make([]int, len(arr))

		copy(arrCopy, arr)
		ch <- arrCopy

		i := 1
		for i < len(arr) {
			if c[i] < i {
				if i%2 == 0 {
					arr[0], arr[i] = arr[i], arr[0]
				} else {
					arr[c[i]], arr[i] = arr[i], arr[c[i]]
				}
				arrCopy = make([]int, len(arr))
				copy(arrCopy, arr)
				ch <- arrCopy
				c[i] += 1
				i = 1
			} else {
				c[i] = 0
				i++
			}
		}
		close(ch)
	}()
	return ch
}

type FullSearch struct{}

func NewFullSearch() *FullSearch {
	return &FullSearch{}
}

func (f *FullSearch) Run(gr *graph.WeightedUndirectedGraph) (*graph.WeightedCycle, error) {
	bestCycle := gr.GetRandomHamiltonian()
	ch := HeapAlgo(gr.GetNodes())
	for arr := range ch {
		path := graph.NewWeightedCycle(gr)
		for _, node := range arr {
			err := path.AddNode(node)
			if err != nil {
				return nil, err
			}
		}
		// path.CalculateWeight()
		// fmt.Println(path)
		if path.CalculateWeight() < bestCycle.CalculateWeight() {
			bestCycle = path
		}
	}
	return bestCycle, nil
}
