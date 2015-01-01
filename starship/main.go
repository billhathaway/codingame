package main

import "fmt"
import "math"

//import "os"

/**
 * The code below will read all the game information for you.
 * On each game turn, information will be available on the standard input, you will be sent:
 * -> the total number of visible enemies
 * -> for each enemy, its name and distance from you
 * The system will wait for you to write an enemy name on the standard output.
 * Once you have designated a target:
 * -> the cannon will shoot
 * -> the enemies will move
 * -> new info will be available for you to read on the standard input.
 **/

func main() {
	for {
		// count: The number of current enemy ships within range
		var count int
		fmt.Scan(&count)
		var closestDist int64 = math.MaxInt64
		var closestEnemy string
		for i := 0; i < count; i++ {
			var enemy string
			var dist int64
			fmt.Scan(&enemy, &dist)
			if dist < closestDistance {
				closestDist = dist
				closestEnemy = enemy
			}
		}
		fmt.Println(closestEnemy) // The name of the most threatening enemy (HotDroid is just one example)
	}
}
