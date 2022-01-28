package hello

import (
	"testing"
)

type PhraseResult struct {
	i        int
	expected string
}

var phraseResult = []PhraseResult{
	{0, "I can eat glass and it doesn't hurt me."},
	{1, "Don't communicate by sharing memory, share memory by communicating."},
	{2, "Hello, world."},
	{4, "If a program is too slow, it must have a loop."},
}

func TestPithyArray(t *testing.T) {
	for i, test := range phraseResult {
		if result := pithySay(i); result != test.expected {
			t.Errorf("greetings() = %q, want %q", result, phraseResult)
		}
	}
}
