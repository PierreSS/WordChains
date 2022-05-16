package wordchains

import (
	"strings"
)

// Same algorithm as simpleWordChains but add a condition to reset the generatedWord to the first word if it encounters a dead end
func (w *Words) MediumWordChains(firstWord, finalWord string) string {
	var generatedWord []byte
	var pos int
	var output []string

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
	}
	output = append([]string{firstWord}, output...)

	return strings.Join(output, "\n")
}
