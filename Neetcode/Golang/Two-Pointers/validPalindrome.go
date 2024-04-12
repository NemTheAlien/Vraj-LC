package main

import "unicode"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	arr := []rune(s)
	for l < r {
		left := unicode.ToLower(arr[l])
		right := unicode.ToLower(arr[r])

		if !isLetterOrDigit(left) {
			l++
			continue
		}
		if !isLetterOrDigit(right) {
			r--
			continue
		}
		if left != right {
			return false
		}
		r--
		l++
	}
	return true
}

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
