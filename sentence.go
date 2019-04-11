package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func GenerateSentence(ord int) {
	m := NewModel(ord)

	words := readSentences("datasets/evil.txt", bufio.ScanWords)

	m.TrainSentences(words)

	for i := 0; i < 15; i++ {
		fmt.Println(m.GenerateSentence())
	}
}

func (m *Model) TrainSentences(words []string) {
	for i := 0; i < len(words)-m.ord; i++ {
		pres := []string{}
		for j := 0; j < m.ord; j++ {
			pres = append(pres, words[i+j])
		}
		pre := strings.Join(pres, " ")
		suf := words[i+m.ord]
		m.Extend(pre, suf)
	}
}

func (m *Model) GenerateSentence() string {

	out := []string{}

	var prefix string = " "
	for !unicode.IsUpper([]rune(prefix)[0]) {
		prefix = randomString(m.keys)
	}

	out = append(out, prefix)

	for {
		suffix := randomString(m.data[prefix])
		out = append(out, suffix)

		if strings.HasSuffix(suffix, ".") {
			break
		}

		prefixes := strings.Split(prefix, " ")

		last := prefixes[len(prefixes)-m.ord+1:]

		prefix = strings.TrimSpace(
			strings.Join(last, " ") + " " + suffix,
		)
	}
	return strings.Join(out, " ")
}

func readSentences(path string, split bufio.SplitFunc) []string {
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
		words = append(words, in)
	}
	return words
}
