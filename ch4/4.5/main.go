package main

import "fmt"

func dedup(items []string) []string {
	i := 0
	last := ""; 
	for _, s := range items {
		if s != last {
			items[i] = s
			last = s
			i++
		}
	}
	return items[:i]
}

func main() {
	s := []string{"a", "a", "b", "a", "c", "b", "b", "b"}
	fmt.Println(dedup(s))
}
