// scrabble2 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

// score calculates the point value for a word
func score(word string) int {
	var total int
	for i := range word {
		switch word[i] {
		case 'e', 'a', 'i', 'o', 'n', 'r', 't', 'l', 's', 'u':
			total += 1
		case 'd', 'g':
			total += 2
		case 'b', 'c', 'm', 'p':
			total += 3
		case 'f', 'h', 'v', 'w', 'y':
			total += 4
		case 'k':
			total += 5
		case 'j', 'x':
			total += 8
		case 'q', 'z':
			total += 10
		}
	}
	return total
}

// copyMap copies an original map to a new one
func copyMap(orig map[uint8]int) map[uint8]int {
	clone := make(map[uint8]int, len(orig))
	for k, v := range orig {
		clone[k] = v
	}
	return clone
}

// checkWords all the words and returns the highest scoring/earliest matching that can be created with tiles
func checkWords(tiles string, dictionary []string) string {
	lc := make(map[uint8]int)
	for i := range tiles {
		lc[tiles[i]]++
	}
	bestScore := 0
	bestWord := ""
	for _, candidate := range dictionary {
		letters := copyMap(lc)
		ok := true
		for j := range candidate {
			if val := letters[candidate[j]]; val > 0 {
				letters[candidate[j]] = val - 1
			} else {
				ok = false
				break
			}
		}
		if ok {
			candidateScore := score(candidate)
			if candidateScore > bestScore {
				bestScore = candidateScore
				bestWord = candidate
			}
		}
	}
	return bestWord
}

// loadData reads from the data file, putting words in dictionary and letters to check in tiles
func loadData() (string, []string) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	var n int
	fmt.Sscan(s.Text(), &n)
	dictionary := make([]string, n)
	i := 0
	for i < n && s.Scan() {
		dictionary[i] = s.Text()
		i++
	}
	fmt.Fprintln(os.Stderr, "Loaded", i, "entries")
	var tiles string
	if s.Scan() {
		tiles = s.Text()
	} else {
		fmt.Fprintln(os.Stderr, "problem parsing file"+s.Err().Error())
		os.Exit(1)
	}
	return tiles, dictionary
}

func main() {
	tiles, dictionary := loadData()
	fmt.Println(checkWords(tiles, dictionary))
}
