package wordchains

import (
	"fmt"
	"strings"
)

// Simple algorithm that run through the firstword letters replacing it by the finalword letter at the same index if not the same then check it's a valid english word and rerun until firstword is equal to finalword
func (w *Words) SimpleWordChains(firstWord, finalWord string) string {
	var generatedWord []byte
	var pos int
	var output []string

	output = append(output, firstWord)
	generatedWord = []byte(firstWord)

	for string(generatedWord) != finalWord {
		// If one letter from final word is different from letter of first word
		if generatedWord[pos] != finalWord[pos] {
			//isEnglishWord = []byte(generatedWord)

			// replace first word with letter from second word at the same index
			copy(generatedWord[pos:], string(finalWord[pos]))

			// if it makes a real english word
			if _, ok := (*w)[string(generatedWord)]; ok {
				output = append(output, string(generatedWord))

				// Set back pos to 0 to check for a new english word or return if generatedWord is finalword
				pos = 0
				continue
			}

			// replace letter from firstword if word is not an english word
			copy(generatedWord[pos:], string(firstWord[pos]))
		}
		pos++

		// Problems : this algorithm stop working if the generatedWord can't find an english word with any letters of finalWord
		if pos >= len(generatedWord) {
			fmt.Println("no english word found, can't continue")

			return ""
		}
	}

	return strings.Join(output, "\n")
}
