package PancakeSorter

import (
	"../InputValidator"
	"bytes"
	"errors"
	"strings"
)

func SortRecursive(pancakeString string, attempts int) (string, int, error) {
	if (! InputValidator.IsValidPancakeString(pancakeString)) {
		return "", 0, errors.New("Invalid pancake string");
	}

	if (strings.Contains(pancakeString, "-")) {
		pancakeString := doSort(pancakeString);
		attempts++;
		return SortRecursive(pancakeString, attempts);
	}

	return pancakeString, attempts, nil;
}

func doSort(pancakeString string) string {
	var reversedPancake = ReverseString(pancakeString);
	var firstNegativeCake = strings.Index(reversedPancake, "-");
	return ReverseString(replaceAtIndex(firstNegativeCake, reversedPancake, swapPlusAndMinus(reversedPancake[firstNegativeCake:])))
}

func ReverseString(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func swapPlusAndMinus(toSwap string) (result string) {
	var b bytes.Buffer

	for i := 0; i < len(toSwap); i++ {
		if string(toSwap[i]) == "-" {
			b.WriteString("+")
		} else {
			b.WriteString("-")
		}
	}
	return b.String();
}

func replaceAtIndex(index int, fullString string, stringToReplace string) string {
	var b bytes.Buffer

	for i := 0; i < len(fullString); i++ {
		if (i >= index) {
			b.WriteString(string(stringToReplace[i - index]))
		} else {
			b.WriteString(string(fullString[i]))
		}
	}

	return b.String()
}