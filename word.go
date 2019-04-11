package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func (m *Model) TrainWords(words []string) {
	for _, word := range words {
		ws := strings.Repeat(" ", m.ord)
		in := []rune(ws + word + ws)

		for i := 0; i < len(in)-m.ord; i++ {
			prefix := string(in[i : i+m.ord])
			suffix := string(in[i+m.ord : i+m.ord+1])
			m.Extend(prefix, suffix)
		}
	}
}

func (m *Model) GenerateWord() string {
	var output strings.Builder

	// Find valid sequence prefix.
	var prefix string
	for !strings.HasPrefix(prefix, " ") {
		prefix = randomString(m.keys)
	}

	output.WriteString(prefix)

	for {
		suffix := randomString(m.data[prefix])
		output.WriteString(suffix)
		if strings.HasSuffix(suffix, " ") {
			break
		}
		n := utf8.RuneCountInString(prefix)
		last := string([]rune(prefix)[n-m.ord+1:])
		prefix = last + suffix
	}

	return strings.TrimSpace(output.String())
}

func readFile(path string, split bufio.SplitFunc) []string {
	words := []string{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(split)
	for scanner.Scan() {
		in := scanner.Text()
		if !utf8.ValidString(in) {
			panic("input contains invalid UTF-8 string")
		}
		if isAlpha(in) {
			words = append(words, in)
		}
	}
	return words
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
