package main

import (
	"fmt"
	"time"
)

type Main interface {
	DoingWork(time time.Duration) error
}

type NotMain interface {
	WatchMovie() error
	Eating() error
}

type Adapter struct {
	deals NotMain
}

type Worker struct{}

func (w *Worker) WatchMovie() error {
	time.Sleep(5 * time.Second)
	return nil
}

func (w *Worker) Eating() error {
	time.Sleep(4 * time.Second)
	return nil
}

func (a *Adapter) DoingWork(time time.Duration) error {
	a.deals.WatchMovie()
	a.deals.Eating()
	return nil
}

func main() {
	fmt.Println("The adapter realisation in code. You can see if you want)")
}
