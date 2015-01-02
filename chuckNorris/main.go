package main

import "fmt"
import "os"
import "bufio"

import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func emit(l uint8, repeat int) {
	//fmt.Fprintf(os.Stderr,"l=%c repeat=%d\n",l,repeat)
	if l == '1' {
		fmt.Print("0 ")
	} else {
		fmt.Print("00 ")
	}
	for i := 0; i <= repeat; i++ {
		fmt.Print("0")
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	MESSAGE := scanner.Text()
	fmt.Fprintln(os.Stderr, MESSAGE)
	var binary string
	for i := 0; i < len(MESSAGE); i++ {
		tmp := strconv.FormatInt(int64(MESSAGE[i]), 2)
		for len(tmp) < 7 {
			tmp = "0" + tmp
		}
		binary += tmp
	}
	fmt.Fprintln(os.Stderr, binary)
	repeat := 0
	current := binary[0]
	for i := 1; i < len(binary); i++ {
		if binary[i] == current {
			repeat++
		} else {
			emit(current, repeat)
			fmt.Print(" ")
			current = binary[i]
			repeat = 0
		}
	}
	emit(current, repeat)
	fmt.Println()
}
