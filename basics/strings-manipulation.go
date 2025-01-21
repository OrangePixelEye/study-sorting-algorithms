package basics

import (
	"regexp"
	"strings"
)

func RemoveLetter(phrase, letter string) {

}

func RemoveAllSymbols(sen string) []string {
	reg, err := regexp.Compile("[^a-zA-Z ]")
	if err != nil {
		return []string{}
	}
	sen = reg.ReplaceAllString(sen, "")
	arr := strings.Split(sen, " ")
	return arr
}
