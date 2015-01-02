package main

import "fmt"
import "sort"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	strengths := make(sort.IntSlice, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &strengths[i])
	}
	// start closest with value greater than possible values
	closest := 10000001
	sort.Sort(strengths)
	for i := 0; i < n-1; i++ {
		delta := strengths[i+1] - strengths[i]
		if delta < closest {
			closest = delta
		}
	}
	fmt.Println(closest)
}
