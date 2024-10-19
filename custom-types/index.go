package main

import "fmt"

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	res := make(map[uint8]int)

	for _, v := range pl {
		_, exists := res[v.Age]

		if !exists {
			res[v.Age] = 1
		} else {
			res[v.Age] += 1
		}
	}

	return res
}

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
		{Age: 34},
		{Age: 13},
		{Age: 18},
	}

	fmt.Print(pl.GetAgePopularity())
}
