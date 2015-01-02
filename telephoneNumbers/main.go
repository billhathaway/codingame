package main

import "fmt"

// total count of objects
var count int

type digitNode struct {
	children map[rune]*digitNode
}

// addTelephone is given a telephone number and updates the count
func (d *digitNode) addTelephone(telephone string) {
	current := d
	for _, digit := range telephone {
		child, exists := current.children[digit]
		if exists {
			current = child
			continue
		}
		current.children[digit] = &digitNode{make(map[rune]*digitNode)}
		count++
		current = current.children[digit]
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	t := &digitNode{make(map[rune]*digitNode)}

	for i := 0; i < N; i++ {
		var telephone string
		fmt.Scan(&telephone)
		t.addTelephone(telephone)
	}

	fmt.Println(count) // The number of elements (referencing a number) stored in the structure.
}
