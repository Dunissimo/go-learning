package main

import "fmt"

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta <= 0 {
		delta = 1
	}

	c.Value = c.Value + delta
}

func (c *Counter) Dec(delta int) {
	if delta <= 0 {
		delta = 1
	}

	c.Value = Max(c.Value-delta, 0)
}

func main() {
	c := Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	fmt.Println(c.Value)

	c.Dec(0)
	c.Dec(30)
	fmt.Println(c.Value)

	c.Dec(100)
	fmt.Println(c.Value)
}
