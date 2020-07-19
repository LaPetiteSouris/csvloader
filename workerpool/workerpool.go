package workerpool

import (
	"math/rand"
	"sync"
)

// Worker is the work horse
type Worker interface {
	ExecuteTask([]string, *sync.WaitGroup, ...interface{}) error
}

// WorkerPool struct
type WorkerPool struct {
	Pool []Worker
	Wg   *sync.WaitGroup
}

// ExecuteJob execute a given task
func (wp *WorkerPool) ExecuteJob(records []string, args ...interface{}) error {
	// look into the pool
	// take out one worker goroutine
	worker := wp.Pool[rand.Intn(len(wp.Pool))]
	// increase waitgroup
	wp.Wg.Add(1)
	// execute job
	go worker.ExecuteTask(records, wp.Wg, args...)
	return nil
}
