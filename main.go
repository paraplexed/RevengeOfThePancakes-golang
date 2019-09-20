package main

import (
	"./PancakeSorter"
	"fmt"
	"os"
)

func main() {
	// Dropping the testcase number on the floor on purpose
	// Not going to bother making this one take it into consideration
	// See c# solution for a more complete handling of it
	var args = os.Args[2:]

	for i := 0; i < len(args); i++ {
		var _, attempts, err = PancakeSorter.SortRecursive(args[i], 0)
		if err != nil {
			fmt.Printf(`There was an error sorting: %s\r\n`, err)
		} else {
			fmt.Printf(`TestCase #%d: %d`, i+1, attempts)
			fmt.Println()
		}
	}
}