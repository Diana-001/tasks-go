package words

import (
	"strings"
)

func findPolyndrome(word1, word2 string) bool {
	strings.ReplaceAll(word1, " ", "")
	strings.ReplaceAll(word2, " ", "")
	len1 := len(word1)
	len2 := len(word2)

	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if word1[i] == word2[len1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
