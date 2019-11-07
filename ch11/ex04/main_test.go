// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomNonPalindrome(rng *rand.Rand) string {
	used := map[rune]struct{}{}
	n := rng.Intn(23) + 2 // 2 ~ 24
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		for {
			r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
			if _, ok := used[r]; !ok {
				runes[i] = r
				used[r] = struct{}{}
				break
			}
		}
	}
	return "a" + string(runes) + "b"
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		if i%3 == 0 {
			switch rng.Intn(5) {
			case 0:
				r = ','
			case 1:
				r = '.'
			case 2:
				r = ' '
			case 3:
				r = '!'
			case 4:
				r = '?'
			}
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) != false {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
