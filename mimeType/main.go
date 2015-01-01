package main

import "fmt"
import "os"
import "bufio"
import "strings"

//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	// Q: Number Q of file names to be analyzed.
	var Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)
	mimeTypes := make(map[string]string, N)
	for i := 0; i < N; i++ {
		// EXT: file extension
		// MT: MIME type.
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &EXT, &MT)
		mimeTypes[strings.ToLower(EXT)] = MT
	}
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := scanner.Text() // One file name per line.
		response := "UNKNOWN"
		lastPeriod := strings.LastIndex(FNAME, ".")
		if lastPeriod > -1 {
			if lastPeriod == (len(FNAME) - 1) {
				fmt.Println(response)
				continue
			}
			extension := strings.ToLower(FNAME[lastPeriod+1:])
			mt, found := mimeTypes[extension]
			if found {
				response = mt
			}
		}
		fmt.Println(response)
	}
}
