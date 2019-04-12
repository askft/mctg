package main

// EditDistance calculates the Levenshtein distance between two strings.
//
// The algorithm is based on
// https://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Levenshtein_distance#C.
func EditDistance(s1, s2 string) int {
	a := []rune(s1)
	b := []rune(s2)

	col := make([]int, len(a)+1)

	for y := 1; y <= len(a); y++ {
		col[y] = y
	}

	for x := 1; x <= len(b); x++ {
		col[0] = x
		lastDiag := x - 1
		for y := 1; y <= len(a); y++ {
			oldDiag := col[y]
			c := 1
			if a[y-1] == b[x-1] {
				c = 0
			}
			col[y] = min3(col[y]+1, col[y-1]+1, lastDiag+c)
			lastDiag = oldDiag
		}
	}

	return col[len(a)]
}

func min3(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}
