package InputValidator

import "regexp"

func IsValidPancakeString(pancakeString string) bool {
	if (pancakeString != "" && len(pancakeString) <= 100) {
		matched, _ := regexp.MatchString(`^[+-]+$`, pancakeString);
		return matched;
	}

	return false
}