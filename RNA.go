package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func transformIntoRNA(dna string)(string) {
	return strings.ReplaceAll(dna, "T", "U")
}

func main() {
	file := os.Args[1]
	buf, err := ioutil.ReadFile(file)

	if (err != nil) {
		fmt.Printf("Error occured while reading input file %v", err)
		return
	}

	dna := string(buf)
	fmt.Println(transformIntoRNA(dna))
}

