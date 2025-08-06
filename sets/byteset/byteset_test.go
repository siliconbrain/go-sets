package byteset

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/siliconbrain/go-sets/sets"
	"github.com/siliconbrain/go-sets/sets/arrayset"
	"github.com/siliconbrain/go-sets/sets/hashset"
	"github.com/stretchr/testify/assert"
)

func TestByteSet(t *testing.T) {
	primeBytes := [...]byte{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
		233, 239, 241, 251,
	}
	set := FromValues(primeBytes[:]...)
	assert.Equal(t, len(primeBytes), set.Len())
	assert.ElementsMatch(t, primeBytes, slices.Collect(set.Values))

	for i := range 256 {
		b := byte(i)
		expected := slices.Contains(primeBytes[:], b)
		adverb := ""
		if !expected {
			adverb = " not"
		}
		assert.Equal(t, expected, set.Contains(b), "byte set should%s contain %d", adverb, b)
	}
}

func BenchmarkByteSet(b *testing.B) {
	benchmarkSet(b, FromValues)
}

func BenchmarkArraySet(b *testing.B) {
	benchmarkSet(b, arrayset.FromValues)
}

func BenchmarkHashSet(b *testing.B) {
	benchmarkSet(b, hashset.FromValues)
}

func benchmarkSet[Set sets.SetOf[byte]](b *testing.B, fromBytes func(...byte) Set) {
	b.Helper()
	for b.Loop() {
		b.StopTimer()
		s := fromBytes(randomGenerateBytes()...)

		b.StartTimer()
		checkBytes(s)
	}
}

func randomGenerateBytes() (bytes []byte) {
	bytes = make([]byte, 100)
	for i := range 100 {
		bytes[i] = byte(rand.IntN(256))
	}
	return
}

func checkBytes(s sets.SetOf[byte]) {
	for i := range 256 {
		_ = s.Contains(byte(i))
	}
}
