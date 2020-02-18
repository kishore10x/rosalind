package main

import (
	"fmt"
	"os"
	"strconv"
)

func modFib(n, k int64) int64 {
	switch(n) {
	case 1: return 1
	case 2: return 1
	default : return modFib(n-2, k)*k + modFib(n-1, k)
	}
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if (err != nil ) {
		panic(err)
	}

	k, err := strconv.ParseInt(os.Args[2], 10, 64)
	if (err != nil ) {
		panic(err)
	}

	fmt.Println(modFib(n, k))
}
