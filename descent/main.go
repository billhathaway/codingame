package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	leftToRight := true
	var fired bool
	for {
		var SX, SY int
		fmt.Scan(&SX, &SY)

		// hh is highest height
		var hh int
		// hm is highest mountain
		var hm int
		for i := 0; i < 8; i++ {
			// MH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
			var MH int
			fmt.Scan(&MH)
			if !fired && MH > hh {
				hm = i
				hh = MH
			}
		}
		if hm == SX && !fired {
			fmt.Println("FIRE")
		} else {
			fmt.Println("HOLD")
		}
		leftToRight = !leftToRight
		fired = false
	}
}
