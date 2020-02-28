package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// ================== Worker ==================

type Worker struct {
	id int
	wg *sync.WaitGroup
}

func NewWorker(id int, wg *sync.WaitGroup) *Worker {
	return &Worker{id, wg}
}

func (w Worker) Start() {
	startTs := rand.Intn(15000)
	time.Sleep(time.Duration(startTs) * time.Millisecond)
	fmt.Printf("Worker #%d started\n", w.id)
	time.Sleep(time.Second)
}

func (w Worker) Work() {
	time.Sleep(300 * time.Millisecond) // Delay where the cap is still in the basket
	durationMs := 2000 + rand.Intn(6000)
	fmt.Printf("Worker #%d will work for %v\n", w.id, float64(durationMs))
	backDurationMs := 3000
	time.Sleep(time.Duration(durationMs+backDurationMs) * time.Millisecond)
}

func (w Worker) Stop() {
	fmt.Printf("Worker #%d stoped\n", w.id)
	w.wg.Done()
}

// ================== Semaphore ==================
// Semaphore is a counter that's used to limit concurrent access to a shared
// resource. Two ways of using a semaphore in Go:
// 1. Incrementing
// 2. Decrementing

func TestConcurrency(t *testing.T) {
	capacity := 2
	maxWorkers := 5
	wg := sync.WaitGroup{}

	semCh := make(chan struct{}, capacity)

	for i := 0; i < capacity; i++ {
		semCh <- struct{}{}
	}

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		w := NewWorker(i, &wg)

		go func() {
			w.Start()
			<-semCh
			fmt.Printf("Worker #%d got assigned work\n", w.id)
			w.Work()
			semCh <- struct{}{}
			fmt.Printf("Worker #%d finished work\n", w.id)
			w.Stop()
		}()
	}

	wg.Wait()
	fmt.Println("All workers are finished!")
}
