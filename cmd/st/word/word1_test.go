package word_test

import (
	. "BookStudy/cmd/st/word"
	"testing"
)

func TestPalindrom(t *testing.T) {
	if !IsPalindrom("detartrated") {
		t.Error(`IsPalindrom("detartrated") = false`)
	}
	if !IsPalindrom("kayak") {
		t.Error(`IsPalindrom("kayak") = false`)
	}
}

func TestNonPalindrom(t *testing.T) {
	if IsPalindrom("palindrom") {
		t.Error(`palindrom("palindrom") = true`)
	}
}
