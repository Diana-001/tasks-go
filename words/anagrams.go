package words

import (
	"reflect"
	"strings"
)

func FindAnagrams() bool {
	s := "diane"
	p := "enaidfchvjb"

	strings.ReplaceAll(s, " ", "")
	strings.ReplaceAll(p, " ", "")

	if len(s) != len(p) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		map1[rune(s[i])]++
		map2[rune(p[i])]++
	}

	if reflect.DeepEqual(map1, map2) {
		return true
	}

	return false
}
