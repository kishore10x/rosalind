package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func reverseComplement(dna string)(string) {
	complement := make([]int32, len(dna))

	ctr := 1

	for _,rune := range(dna) {
		complement[len(dna)-ctr] = getComplement(rune)
		ctr = ctr + 1
	}

	return string(complement)
}

func getComplement(b int32) int32{
	switch b {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	default:
		return '0'
	}
}

func main() {
	file := os.Args[1]
	buf, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("Error occured while reading input file %v", err)
		return
	}

	dna := string(buf)
	fmt.Println(reverseComplement(dna))
}

