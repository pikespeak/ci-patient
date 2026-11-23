// Package wordstat counts words in a text.
package wordstat

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
)

// Entry is one word and how often it occurs.
type Entry struct {
	Word  string
	Count int
}

// Top returns the n most frequent words, ties in alphabetical order.
func Top(r io.Reader, n int) ([]Entry, error) {
	if n < 1 {
		return nil, fmt.Errorf("n must be at least 1, got %d", n)
	}
	counts := map[string]int{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		for _, w := range strings.FieldsFunc(strings.ToLower(s.Text()), func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsDigit(c)
		}) {
			counts[w]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("read text: %w", err)
	}
	out := make([]Entry, 0, len(counts))
	for w, c := range counts {
		out = append(out, Entry{w, c})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Word < out[j].Word
	})
	return out[:min(n, len(out))], nil
}
