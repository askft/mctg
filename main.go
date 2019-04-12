package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
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
	)
	flag.Parse()

	if *path == "" {
		fmt.Println("must suppply path (-path <value>)")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	m := NewModel(*ord)

	// bible := readSentences("datasets/bible.txt", bufio.ScanWords)
	// atlas := readSentences("datasets/atlas-shrugged.txt", bufio.ScanWords)
	// m.TrainSentences(bible)
	// m.TrainSentences(atlas)
	// for i := 0; i < *num; i++ {
	// 	fmt.Println(m.GenerateSentences(2))
	// 	fmt.Println()
	// }

	words := readFile(*path, bufio.ScanLines)
	m.TrainWords(words)
	for i := 0; i < *num; i++ {
		fmt.Println(m.GenerateWord())
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
