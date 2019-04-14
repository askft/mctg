package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
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
		typ  = flag.String("type", "", "'word' or 'sentence'")
	)
	flag.Parse()

	if *path == "" {
		fmt.Println("[warning] path not supplied")
	}

	if *typ == "" {
		fmt.Println("[warning] type not supplied")
	}

	rand.Seed(time.Now().UnixNano())

	m := NewModel(*ord)

	switch *typ {
	case "word":
		generateWords(m, *path, *num)
	case "sentence":
		generateSentences(m, *path, *num)
	}
}

func generateWords(m *Model, path string, num int) {
	words := readFile(path, bufio.ScanLines)
	m.TrainWords(words)
	result := []string{}
	for i := 0; i < num; i++ {
		result = append(result, m.GenerateWord())
	}
	sort.Strings(result)
	for i := 0; i < num; i++ {
		fmt.Println(result[i])
	}
}

func generateSentences(m *Model, path string, num int) {
	bible := readSentences(path, bufio.ScanWords)
	m.TrainSentences(bible)

	for i := 0; i < num; i++ {
		fmt.Println(m.GenerateSentences(2))
		fmt.Println()
	}
}

func (m *Model) Extend(k, v string) {
	if _, ok := m.data[k]; !ok {
		m.data[k] = []string{}
		m.keys = append(m.keys, k)
	}
	m.data[k] = append(m.data[k], v)
}

// fails if len(s) == 0
func randomString(s []string) string {
	return s[rand.Intn(len(s))]
}
