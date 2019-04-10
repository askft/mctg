package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func readWords(path string) []string {
	words := []string{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func main() {
	rand.Seed(time.Now().Unix())

	ord := 2
	N := 100

	words := readWords("alice_oz.txt")

	stuff := map[string][]string{}
	keys := []string{}

	for i := 0; i < len(words)-ord; i++ {

		prefixes := []string{}
		for j := 0; j < ord; j++ {
			prefixes = append(prefixes, words[i+j])
		}

		prefix := strings.Join(prefixes, " ")

		suffix := words[i+ord]

		if _, ok := stuff[prefix]; !ok {
			stuff[prefix] = []string{}
			keys = append(keys, prefix)
		}
		stuff[prefix] = append(stuff[prefix], suffix)
	}

	// for k, v := range stuff {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }

	fmt.Println()

	prefix := keys[rand.Intn(len(keys))]
	// fmt.Println("Random prefix:", prefix)

	output := []string{prefix}

	for i := 0; i < N; i++ {
		// fmt.Println()

		suffixes := stuff[prefix]
		// fmt.Printf("suffixes: %v.\n", suffixes)

		if len(suffixes) == 0 {
			break
		}

		suffix := suffixes[rand.Intn(len(suffixes))]
		// fmt.Printf("... chose %s.\n", suffix)

		output = append(output, suffix)

		prefixes := strings.Split(prefix, " ")
		// fmt.Printf("prefixes: %v.\n", prefixes)

		lastOrdPrefixes := prefixes[len(prefixes)-ord+1:]
		// fmt.Println("lastOrdPrefixes:", lastOrdPrefixes)

		prefix = strings.TrimSpace(strings.Join(lastOrdPrefixes, " ") + " " + suffix)
		// fmt.Println("prefix:", prefix)

	}

	fmt.Println(strings.Join(output, " "))
}
