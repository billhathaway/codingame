package main

import "fmt"

var L, C, N int

type rideFill struct {
	count     int
	nextGroup int
}

var cached = make(map[int]rideFill)
var groups = make([]int, 0)

func cacheRideFill(startGroup int) rideFill {
	rf, ok := cached[startGroup]
	if ok {
		return rf
	}
	groupCount := 0
	for rf.count+groups[(startGroup+groupCount)%N] <= L && groupCount < N {
		rf.count += groups[(startGroup+groupCount)%N]
		groupCount++
	}
	rf.nextGroup = (startGroup + groupCount) % N
	cached[startGroup] = rf
	return rf
}

func main() {
	fmt.Scan(&L, &C, &N)

	var g int
	for i := 0; i < N; i++ {
		fmt.Scan(&g)
		groups = append(groups, g)
	}

	rf := cacheRideFill(0)
	totalPay := rf.count
	for i := 1; i < C; i++ {
		rf = cacheRideFill(rf.nextGroup)
		totalPay += rf.count
	}
	fmt.Println(totalPay)
}

