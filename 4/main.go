package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Workers(wg *sync.WaitGroup, tsk chan interface{}, quit chan bool, n int) {
	wg.Add(1)
	var work interface{}
	for {
		//rand.Seed(time.Now().Unix())
		select {
		case work = <-tsk:
			log.Println(n, work)
			time.Sleep(3 * time.Second)
		case <-quit:
			log.Printf("Worker %d stop working\n", n)
			wg.Done()
			return
		}
	}

}
func main() {
	var n int
	task := make([]interface{}, 0)
	task = append(task,
		"task1",
		"task2",
		"task3",
		"task4",
	)
	var wg sync.WaitGroup
	cancel := make(chan bool)
	tsk := make(chan interface{})
	sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan,
	// 	syscall.SIGHUP,
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("How many workers do u want???")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		go Workers(&wg, tsk, cancel, i)
	}
	for {
		select {
		case <-sigChan:
			close(cancel)
			log.Println("stop")
			wg.Wait()
			return
		default:
			rand.Seed(time.Now().UnixNano())
			tsk <- task[rand.Intn(4)]
		}
	}
}
