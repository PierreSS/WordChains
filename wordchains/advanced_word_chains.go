package wordchains

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Same algorithm as simpleWordChains but add a condition to reset the generatedWord to the first word if it encounters a dead end
func (w *Words) AdvancedWordChains(firstWord, finalWord, firstWordEver string) string {
	var generatedWord []byte
	var pos int
	var output []string

	deadline := time.Now().Add(1 * time.Second)
	generatedWord = []byte(firstWord)

	for string(generatedWord) != finalWord {
		// If one letter from final word is different from letter of first word
		if generatedWord[pos] != finalWord[pos] {
			//isEnglishWord = []byte(generatedWord)

			// replace first word with letter from second word at the same index
			copy(generatedWord[pos:], string(finalWord[pos]))

			// if it makes a real english word and the generated word isn't a dead end
			if _, ok := (*w)[string(generatedWord)]; ok && !(*w)[string(generatedWord)] {
				output = append(output, string(generatedWord))

				// Set back pos to 0 to check for a new english word or return if generatedWord is finalword
				pos = 0
				continue
			}

			// replace letter from firstword if word is not an english word
			copy(generatedWord[pos:], string(firstWord[pos]))
		}
		pos++

		// Reset and omit the last generatedWord (dead end))
		if pos >= len(generatedWord) {
			(*w)[string(generatedWord)] = true
			generatedWord = []byte(firstWord)
			pos = 0
			output = nil
		}

		// After 1 second of not finding any pattern, change one letter randomly of first word one by one until it's not a dead end
		if time.Now().After(deadline) {
			for k := range *w {
				for i := range []byte(firstWordEver) {
					expression := []byte(firstWordEver)
					expression[i] = '.'

					match, err := regexp.MatchString("^"+string(expression)+"$", k)
					if err != nil {
						fmt.Println("error matching map keys with expression")
						return ""
					}

					if match && !(*w)[k] {
						(*w)[firstWord] = true

						return w.AdvancedWordChains(k, finalWord, firstWordEver)
					}
				}
			}

		}
	}

	output = append([]string{firstWord}, output...)

	return firstWordEver + "\n" + strings.Join(output, "\n")
}
