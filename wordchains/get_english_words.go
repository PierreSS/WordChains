package main

import (
	"bufio"
	"os"
)

// Retrieve all words from english_words file and return as a map string of string
func (w *Words) GetEnglishWords() error {
	file, err := os.Open("english_words")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		(*w)[scanner.Text()] = false
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
