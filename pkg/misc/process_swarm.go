package misc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ProcessSwarm creates a worker pool
func ProcessSwarm() {
	fmt.Println("Beginning Process Swarm")
	var wg sync.WaitGroup

	commands := buildCommands(420)
	getTask := taskGenerator(commands, 420)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(&wg, getTask)
	}

	wg.Wait()
	fmt.Println("Process Swarm Complete")
}

func buildCommands(count int) *[]string {
	commands := make([]string, count)
	for i := 0; i < count; i++ {
		commands[i] = fmt.Sprintf("COMMAND: %d", i)
	}
	return &commands
}

func taskGenerator(commands *[]string, count int) func() (uint32, func()) {
	return func() (uint32, func()) {
		command := (*commands)[randRange(0, count-1)]
		task := func() {
			fmt.Println(command)
			time.Sleep(time.Duration(randRange(1, 5)) * time.Second)
		}
		return rand.Uint32(), task
	}

}

func worker(wg *sync.WaitGroup, getTask func() (uint32, func())) {
	tag := randRange(1, 512*2*2*2*2*2*2)
	generator := genRange(1, 10)
	exitCode := generator()

	for {
		_, task := getTask()
		taskID := generator()

		if taskID == exitCode {
			break
		}

		fmt.Printf("Worker: %d Running Task\n", tag)
		task()
	}

	fmt.Printf("Worker: %d Finished\n", tag)
	wg.Done()
}

func genRange(min int, max int) func() int {
	return func() int {
		return randRange(min, max)
	}
}

func randRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
