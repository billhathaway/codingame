package main

import "fmt"

func main() {
	// LX: the X position of the light of power
	// LY: the Y position of the light of power
	// TX: Thor's starting X position
	// TY: Thor's starting Y position
	var LX, LY, TX, TY int
	fmt.Scan(&LX, &LY, &TX, &TY)

	for {
		// E: The level of Thor's remaining energy, representing the number of moves he can still make.
		var E int
		fmt.Scan(&E)
		var move string
		switch {
		case TY < LY:
			move = "S"
			TY++
		case TY > LY:
			move = "N"
			TY--
		}
		switch {
		case TX < LX:
			move += "E"
			TX++
		case TX > LX:
			move += "W"
			TX--
		}

		fmt.Println(move) // A single line providing the move to be made: N NE E SE S SW W or NW
	}
}
