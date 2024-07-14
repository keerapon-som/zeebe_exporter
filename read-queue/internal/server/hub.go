package server

import (
	"fmt"
	"readq/internal/service"
	"time"
)

func Run(pipe <-chan []byte) {
	workerPipe := make(chan []byte, 1000)

	InitWorker(workerPipe, 5)

	saveticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case data := <-pipe:
			workerPipe <- data
		case <-saveticker.C:
			service.PerformBatchRecord()
		}
	}

}

func InitWorker(msg <-chan []byte, numberOfWorkers int) {
	for i := 0; i < numberOfWorkers; i++ {
		go worker(msg, i)
	}
}

func worker(msg <-chan []byte, id int) {
	fmt.Println("Worker", id, "started")
	// saveticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case data := <-msg:
			service.TypeClasify(data)
		}
	}
}

// func run() {
// 	msg := make(chan []byte)
// 	InitWorker(5, msg)
// }
