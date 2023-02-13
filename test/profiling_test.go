package test

import (
	"testing"

	index "falconEmailBackend/scripts"
)

// BenchmarkIndexer test for indexer() function
func BenchmarkIndexer(b *testing.B) {
	index.Indexer()
}
