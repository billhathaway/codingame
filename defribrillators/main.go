package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type defibrillator struct {
	name      string
	address   string
	latitude  float64
	longitude float64
}

var defibrillators = make([]defibrillator, 0)

func findClosest(longitude float64, latitude float64) string {
	var closestDistance float64 = math.MaxFloat64
	var closest string = "NOTFOUND"
	for _, defib := range defibrillators {
		x := (defib.longitude - longitude) * math.Cos((latitude+defib.latitude)/2)
		y := defib.latitude - latitude
		d := math.Sqrt(x*x+y*y) * 6371
		if d < closestDistance {
			closest = defib.name
			closestDistance = d
			fmt.Fprintf(os.Stderr, "Setting closest d=%f name=%s\n", d, defib.name)
		}
	}
	return closest
}

// parseFrFloat parses strings in the format of French floating point representations
// they have a comma instead of a period
// we don't need to do error handling besides writing a message since that will cause the test to fail
func parseFrFloat(data string) float64 {
	parsed, err := strconv.ParseFloat(strings.Replace(data, ",", ".", 1), 64)
	if err != nil {
		fmt.Printf("Could not parse value %s %s\n", data, err.Error())
		return parsed
	}
	return parsed
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var LON string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &LON)

	var LAT string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &LAT)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		components := strings.Split(line, ";")
		defibrillators = append(defibrillators, defibrillator{
			name:      components[1],
			address:   components[2],
			longitude: parseFrFloat(components[4]),
			latitude:  parseFrFloat(components[5]),
		})
	}
	fmt.Println(findClosest(parseFrFloat(LON), parseFrFloat(LAT)))
}
