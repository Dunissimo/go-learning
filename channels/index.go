package main

import "fmt"

func SumWorker(numsCh chan []int, sumCh chan int) {
	s := 0

	for v := range numsCh {
		for _, val := range v {
			s += val
		}
	}

	sumCh <- s
}

func main() {
	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)

	numsCh <- nil

	numsCh <- []int{}

	numsCh <- []int{10, 10, 10}

	numsCh <- []int{500, 5, 10, 25}

	close(numsCh)

	result := <-sumCh
	fmt.Println("Result:", result)
}
