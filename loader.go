package main

import (
	"strconv"
	"sync"

	pool "github.com/LaPetiteSouris/csvloader/workerpool"
)

// LoadRecordToDatabase take records and dump to Database
func LoadRecordToDatabase(records []string, numberOfGoroutine int, args ...interface{}) error {
	var wg sync.WaitGroup

	// Initiate worker pool
	// Use the corresponding worker type
	workerArray := make([]pool.Worker, 0)
	for i := 0; i < numberOfGoroutine; i++ {
		w := &pool.PostgresWorker{ID: strconv.FormatInt(int64(i), 10)}
		workerArray = append(workerArray, w)
	}
	workerPool := &pool.WorkerPool{Wg: &wg, Pool: workerArray}
	workerPool.ExecuteJob(records, args...)
	wg.Wait()
	return nil
}
