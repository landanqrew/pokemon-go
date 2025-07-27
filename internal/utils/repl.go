package utils

import "strings"



func cleanInput(text string) []string {
	subStrings := strings.Split(text, " ")
	finalStrings := make([]string, 0)
	for _, str := range subStrings {
		if str != "" {
			finalStrings = append(finalStrings, str)
		}
	}

	return finalStrings
}