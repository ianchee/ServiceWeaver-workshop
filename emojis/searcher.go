package main

import (
	"context"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver"
	"golang.org/x/exp/slices"
)

// Searcher is an emoji search engine component.
type Searcher interface {
	// Search returns the set of emojis that match the provided query.
	Search(context.Context, string) ([]string, error)
}

// searcher is the implementation of the Searcher component.
type searcher struct {
	weaver.Implements[Searcher]
}

func (s *searcher) Search(ctx context.Context, query string) ([]string, error) {
	// Perform the search. First, we lowercase the query and split it into words.

	words := strings.Fields(strings.ToLower(query))
	var results []string
	for emoji, labels := range emojis {
		if matches(labels, words) {
			results = append(results, emoji)
		}
	}
	sort.Strings(results)
	return results, nil

}

func matches(labels, words []string) bool {
	for _, word := range words {
		if !slices.Contains(labels, word) {
			return false
		}
	}
	return true
}
