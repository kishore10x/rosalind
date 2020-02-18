package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var dnaStringID string
	var dnaString string
	var gcContent float32

	var tmpStringID string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sb strings.Builder

	updateIfReqd := func() {
		dnaString = sb.String()
		if (len(dnaString) > 0) {
			tmpPercentage := gcPercentage(dnaString)
			if (tmpPercentage >= gcContent) {
				dnaStringID = tmpStringID
				gcContent = tmpPercentage
			}
		}
	}

	for scanner.Scan() {
		txt := scanner.Text()

		if (strings.HasPrefix(txt, ">")) {
			updateIfReqd()
			sb.Reset()
			tmpStringID =  strings.TrimLeft(txt, ">")
			continue
		}

		sb.WriteString(txt)
	}

	updateIfReqd()

	fmt.Println(dnaStringID)
	fmt.Println(gcContent)
}

func gcPercentage(dna string) float32 {
	gcCount := 0

	for _,rune := range(dna) {
		switch rune {
		case 'C','G':
			gcCount++
		}
	}

	return float32(gcCount)/float32(len(dna)) * 100
}

