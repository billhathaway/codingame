package main

import "fmt"
import "os"
import "bufio"
import "bytes"
import "unicode"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	scanner.Scan()
	T := scanner.Text()

	letters := make([]string, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		letters[i] = scanner.Text()
	}
	outbuf := make([]bytes.Buffer, H)

	for _, char := range T {
		index := 26
		uc := unicode.ToUpper(char)
		if uc >= 'A' && uc <= 'Z' {
			index = int(uc - 'A')
		}
		for i := 0; i < H; i++ {
			offset := index * L
			outbuf[i].WriteString(letters[i][offset : offset+L])
		}
	}
	for i := 0; i < H; i++ {
		fmt.Println(outbuf[i].String())
	}
}
