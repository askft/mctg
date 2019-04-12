package main

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	tables := []struct {
		a, b string
		n    int
	}{
		{"commodore", "commander", 4},
		{"kitten", "sitting", 3},
		{"honey", "honeybee", 3},
		{"kitten", "knitted", 2},
		{"banana", "ananab", 2},
		{"book", "bokop", 2},
		{"book", "back", 2},
		{"x", "x", 0},
		{"x", "", 1},
		{"", "x", 1},
		{"", "", 0},
	}

	for _, table := range tables {
		n := EditDistance(table.a, table.b)
		if n != table.n {
			t.Errorf("EditDistance(%s, %s) was incorrect, got: %d, want: %d",
				table.a, table.b, n, table.n)
		}
	}
}

func TestMin3(t *testing.T) {
	tables := []struct {
		a, b, c, n int
	}{
		{1, 2, 3, 1},
		{3, 1, 2, 1},
		{3, 2, 1, 1},
		{3, 2, 1, 1},
		{1, 3, 2, 1},
		{2, 1, 3, 1},
		{-1, 0, -1, -1},
		{0, 0, 0, 0},
		{-10, -5, -15, -15},
	}

	for _, table := range tables {
		m := min3(table.a, table.b, table.c)
		if m != table.n {
			t.Errorf("min3(%d, %d, %d) was incorrect, got: %d, want: %d.",
				table.a, table.b, table.c, m, table.n)
		}
	}
}
