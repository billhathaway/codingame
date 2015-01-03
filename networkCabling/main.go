package main

import "fmt"
import "math"
import "sort"

type house struct {
	X int64
	Y int64
}

func main() {
	var N int64
	fmt.Scan(&N)

	houses := make([]house, N)

	var minX int64 = math.MaxInt64
	var maxX int64 = math.MinInt64
	var totalY int64
	var X, Y int64
	var i int64
	yOffsets := make(sort.IntSlice, N)
	for i = 0; i < N; i++ {
		fmt.Scan(&X, &Y)
		houses[i] = house{X, Y}
		if X < minX {
			minX = X
		}
		if X > maxX {
			maxX = X
		}
		totalY += Y
		yOffsets[i] = int(Y)
	}
	sort.Sort(yOffsets)
	var yOffset int64 = int64(yOffsets[int(N/2)])

	totalCableLength := maxX - minX
	for _, h := range houses {
		if h.Y > yOffset {
			totalCableLength += (h.Y - yOffset)
		} else {
			totalCableLength += (yOffset - h.Y)
		}
	}
	fmt.Println(totalCableLength)
}
