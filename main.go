package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type Model struct {
	data map[string][]string
	keys []string
	ord  int
}

func NewModel(ord int) *Model {
	return &Model{
		data: map[string][]string{},
		keys: []string{},
		ord:  ord,
	}
}

func main() {
	var (
		ord  = flag.Int("ord", 2, "Markov chain order")
		num  = flag.Int("num", 15, "Number of results to generate")
		path = flag.String("path", "", "Relative path of input")
	)
	flag.Parse()

	if *path == "" {
		fmt.Println("must suppply path (-path <value>)")
		os.Exit(1)
	}

	words := readFile(*path, bufio.ScanLines)

	rand.Seed(time.Now().UnixNano())

	m := NewModel(*ord)

	m.Train(words)

	for i := 0; i < *num; i++ {
		fmt.Println(m.GenerateWord())
	}
}

func (m *Model) Train(words []string) {
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

func (m *Model) Extend(k, v string) {
	if _, ok := m.data[k]; !ok {
		m.data[k] = []string{}
		m.keys = append(m.keys, k)
	}
	m.data[k] = append(m.data[k], v)
}

// func validFirstWord(word string) {
// return if first character is capital letter
// }

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

// ————————————————————————————————————————————————————————
//   Util
//

// fails if len(s) == 0
func randomString(s []string) string {
	return s[rand.Intn(len(s))]
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
