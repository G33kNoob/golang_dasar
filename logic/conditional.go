package logic

import (
	"fmt"
	"strings"
)

func Conditional(words string) {
	for index, word := range words {
		abjad := strings.Contains("aiueo", string(word))
		if abjad && index%2 == 0 {
			fmt.Println("hurufnya adalah =", string(word))
		}

		// fmt.Println("index ke=", index, "letter :", word, string(word))
	}
}
