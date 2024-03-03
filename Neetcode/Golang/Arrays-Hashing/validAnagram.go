package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seenS := make(map[rune]int)

	for _, ch := range s {
		seenS[ch]++
	}
	for _, ch := range t {
		if _, ok := seenS[ch]; ok {
			seenS[ch]--
			if seenS[ch] == 0 {
				delete(seenS, ch)
			}
		}
	}
	return len(seenS) == 0
}
