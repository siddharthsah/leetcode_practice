package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mappingS := make(map[byte]byte)
	mappingT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if _, ok := mappingS[charS]; !ok {
			mappingS[charS] = charT
		} else if mappingS[charS] != charT {
			return false
		}

		if _, ok := mappingT[charT]; !ok {
			mappingT[charT] =charS
		} else if mappingT[charT] != charS {
			return false
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	if isIsomorphic(s, t) {
		fmt.Println("strings are isomophic")
	} else {
		fmt.Println("strings are not isomorphic")
	}
}