//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	firstSign = iota
	secondSign
	text
)

/*
 * Complete the 'camelCase' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func camelCase(s string) {
	signs := strings.Split(s, ";")
	t := signs[text]
	var final string

	reg, _ := regexp.Compile("(([a-z0-9]+)|\\s+)([A-Z])")
	result := reg.ReplaceAllString(t, "$1 $3")
	words := strings.Split(strings.TrimSpace(result), " ")

	if signs[firstSign] == "C" {
		switch signs[secondSign] {
		case "M":
			words[len(words)-1] = strings.Title(words[len(words)-1]) + "()"
		case "C":
			for k, w := range words {
				words[k] = strings.Title(w)
			}
		case "V":
			for i := 1; i < len(words); i++ {
				words[i] = strings.Title(words[i])
			}
		}
		// final all cases are checked successfully
		final = strings.Join(words, "")
	}
	if signs[firstSign] == "S" {
		switch signs[secondSign] {
		case "M":
			words[len(words)-1] = strings.ToLower(strings.Replace(words[len(words)-1], "()", "", 1))
		case "C":
			for k, w := range words {
				words[k] = strings.ToLower(w)
			}
		case "V":
			for k, w := range words {
				words[k] = strings.ToLower(w)
			}
		}
		// final all cases are checked successfully
		final = strings.Join(words, " ")
	}

	fmt.Println(final)
}

func main() {

	var texts []string
	var reader = bufio.NewReader(os.Stdin)

	for true {
		message, err := reader.ReadString('\n')
		if err == nil {
			message = strings.Trim(message, "\n\r")
			texts = append(texts, message)
		} else {
			texts = append(texts, message)
			break
		}
	}

	for _, text := range texts {
		camelCase(text)
	}
}
