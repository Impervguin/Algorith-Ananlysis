package main

import (
	"fmt"
	"io"
	"lab5/internal/loading"
	"lab5/internal/recipedb"
	"lab5/internal/recipeparse"
	"time"
)

type PipelineTimes struct {
	LoadingStart time.Time
	LoadingEnd   time.Time
	ParsingStart time.Time
	ParsingEnd   time.Time
	StorageStart time.Time
	StorageEnd   time.Time
}

type PipelineUnit struct {
	recipedb.Recipe
	PipelineTimes
	PageReader io.ReadCloser
	Err        error
}

var urlId int64 = 0

func getNextUrlId() int64 {
	urlId++
	return urlId
}

func pipelineStarter(urls []string, IssueID int64) <-chan PipelineUnit {
	out := make(chan PipelineUnit)
	go func() {
		for _, url := range urls {
			unit := PipelineUnit{}
			unit.ID = getNextUrlId()
			unit.IssueID = IssueID
			unit.Url = url
			out <- unit
		}
		close(out)
	}()
	return out
}

func loadingStart(in <-chan PipelineUnit) <-chan PipelineUnit {
	out := make(chan PipelineUnit)
	go func() {
		for unit := range in {
			unit.LoadingStart = time.Now()
			unit.PageReader, unit.Err = loading.LoadPage(unit.Url)
			if unit.PageReader == nil {
				unit.Err = fmt.Errorf("failed to load page: %w", unit.Err)
			}
			unit.LoadingEnd = time.Now()
			out <- unit
		}
		close(out)
	}()
	return out
}

func parsingStart(in <-chan PipelineUnit) <-chan PipelineUnit {
	out := make(chan PipelineUnit)
	go func() {
		for unit := range in {
			unit.ParsingStart = time.Now()
			if unit.Err == nil {
				unit.Err = recipeparse.ParseRecipe(unit.PageReader, &unit.Recipe)
				unit.PageReader.Close()
			}
			unit.ParsingEnd = time.Now()
			out <- unit
		}
		close(out)
	}()
	return out
}

func storageStart(in <-chan PipelineUnit, pgs *recipedb.PgsStorage) (<-chan PipelineUnit, error) {
	if pgs == nil {
		return nil, fmt.Errorf("storage is not initialized")
	}
	out := make(chan PipelineUnit)
	go func() {
		for unit := range in {
			unit.StorageStart = time.Now()
			if unit.Err == nil {
				unit.Err = pgs.PutRecipe(&unit.Recipe)
			}
			unit.StorageEnd = time.Now()
			out <- unit
		}
		close(out)
	}()
	return out, nil
}
