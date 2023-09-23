package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const worker = 10

type Person struct {
	sync.Mutex
	Age int
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func Work(wg *sync.WaitGroup, data *Person, name string) {
	data.Lock()
	fmt.Println("-------------------------")
	fmt.Printf("%s take data: age = %d\n", name, data.Age)

	number := rand.Intn(10)
	fmt.Printf("Value: %d\n", number)

	data.SetAge(number)
	fmt.Printf("After %s change data: age = %d\n", name, data.Age)

	time.Sleep(1 * time.Second)
	data.Unlock()

	wg.Done()
}

func main() {
	data := Person{Age: 5}

	var wg sync.WaitGroup
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go Work(&wg, &data, fmt.Sprintf("Worker %d", i))
	}

	wg.Wait()
}
