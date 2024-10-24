package main

import (
	"bufio"
	"context"
	"fmt"
	"lab5/internal/recipedb"
	"os"
	"sort"
	"time"
)

const IssueID int64 = 9161

type ProcessingInterval struct {
	Start time.Duration
	End   time.Duration
	Type  string
	ID    int64
}

func NewProcessingInterval(start, end time.Duration, t string, id int64) *ProcessingInterval {
	return &ProcessingInterval{
		Start: start,
		End:   end,
		Type:  t,
		ID:    id,
	}
}

func (pi *ProcessingInterval) String() string {
	return fmt.Sprintf("%d\t%s\t%f\t%f", pi.ID, pi.Type, pi.Start.Seconds(), pi.End.Seconds())
}

func ReadUrls(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func main() {
	programStartTime := time.Now()
	urls, err := ReadUrls("links.txt")
	if err != nil {
		fmt.Printf("Can't read URLs from file: %v\n", err)
		return
	}

	pgs, err := recipedb.NewPgsStorage(context.TODO())
	if err != nil {
		fmt.Printf("Can't initialize PGS storage: %v\n", err)
		return
	}

	units := pipelineStarter(urls, IssueID)

	unitsPages := loadingStart(units)

	unitsParsed := parsingStart(unitsPages)

	unitsPushed, err := storageStart(unitsParsed, pgs)

	if err != nil {
		fmt.Printf("Can't start storage goroutine: %v\n", err)
		return
	}

	var failed int
	var success int
	// unitsProcessed := make([]PipelineUnit, 0, len(units))
	processingIntervals := make([]ProcessingInterval, 0, 3*len(units))
	for unit := range unitsPushed {
		// if unit.Err != nil {
		// 	failed++
		// 	fmt.Printf("Error processing URL %d: %v\n", unit.ID, unit.Err)
		// } else {
		// 	success++
		// 	fmt.Printf("Successfully processed URL %d\n", unit.ID)
		// }
		// unitsProcessed = append(unitsProcessed, unit)
		processingIntervals = append(processingIntervals, *NewProcessingInterval(unit.LoadingStart.Sub(programStartTime), unit.LoadingEnd.Sub(programStartTime), "loading", unit.ID))
		processingIntervals = append(processingIntervals, *NewProcessingInterval(unit.ParsingStart.Sub(programStartTime), unit.ParsingEnd.Sub(programStartTime), "parsing", unit.ID))
		processingIntervals = append(processingIntervals, *NewProcessingInterval(unit.StorageStart.Sub(programStartTime), unit.StorageEnd.Sub(programStartTime), "storage", unit.ID))
	}

	sort.Slice(processingIntervals, func(i, j int) bool {
		return processingIntervals[i].Start.Seconds() < processingIntervals[j].Start.Seconds()
	})
	fmt.Printf("Processed %d URLs successfully, %d failed\n", success, failed)

	file, err := os.Create("pipeline.log")
	if err != nil {
		fmt.Printf("Can't create pipeline.log file: %v\n", err)
		return
	}
	defer file.Close()

	for _, pi := range processingIntervals {
		file.WriteString(pi.String() + "\n")
	}
	fmt.Printf("Pipeline log written to pipeline.log\n")

}
