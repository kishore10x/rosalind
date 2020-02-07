package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Count(dna string){
	counters := map[int32]int{}

	for _,rune := range(dna) {
		counters[rune] += 1
	}
	
	fmt.Printf( "%d %d %d %d" , counters['A'], counters['C'], counters['G'], counters['T'])
}

func main() {
	file := os.Args[1]
	buf, err := ioutil.ReadFile(file)

	if (err != nil) {
		fmt.Printf("Error occured while reading input file %v", err)
		return
	}

	Count(string(buf))
}
