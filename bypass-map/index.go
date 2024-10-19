package main

import "fmt"

func MostPopularWord(words []string) string {
	max := words[0]
	popular := make(map[string]int)

	for _, w := range words {
		popular[w]++
	}

	for w, count := range popular {
		if count > popular[max] {
			max = w
		}
	}

	return max
}

func main() {
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello"}))
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello", "world", "world"}))
	fmt.Println(MostPopularWord([]string{"one", "two", "three", "four", "five"}))
	fmt.Println(MostPopularWord([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
	fmt.Println(MostPopularWord([]string{"a", "b", "b", "a"}))
}
