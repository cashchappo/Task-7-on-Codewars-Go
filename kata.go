//only words and numbers

package main

import "fmt"

func main() {
	str := "" //you can enter your own string to check
	fmt.Println(alphanumeric(str))
}

func alphanumeric(str string) bool {

	var checker = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	}

	var strnew string

	for i := 0; i != len(str); i++ {
		for i1 := 0; i1 != len(checker); i1++ {
			if string(str[i]) == checker[i1] {
				strnew += checker[i1]
			}
		}
	}

	var result bool

	if len(strnew) == 0 {
		result = false
	} else if strnew == str {
		result = true
	} else {
		result = false
	}

	return result
}
