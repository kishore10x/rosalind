package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var dnaStrings = make([]string,2)

	for i := 0; i<2; i++{
		scanner.Scan()
		dnaStrings[i] = scanner.Text()
	}

	fmt.Println(doCalc(dnaStrings[0], dnaStrings[1]))
}

// s1 & s2 are of equal length
func doCalc(s1, s2 string) int {
	var hd int

	for n, rune := range s1 {
		if (int32(s2[n]) != rune) {
			hd++
		}
	}

	return hd
}
