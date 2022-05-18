package main

import (
	"fmt"
	"os"
)

type Words map[string]bool

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error, input must contains first_word and second_word arguments")
		return
	}

	arg1, arg2 := os.Args[1], os.Args[2]
	if len(arg1) != len(arg2) {
		fmt.Println("Error, the length of the string must be equal")
		return
	}

	w := &Words{}
	if err := w.GetEnglishWords(); err != nil {
		fmt.Println("Error reading words from file : " + err.Error())
		return
	}

	if _, ok := (*w)[arg1]; !ok {
		fmt.Println("First argument is not a valid english word")
		return
	}
	if _, ok := (*w)[arg2]; !ok {
		fmt.Println("Second argument is not a valid english word")
		return
	}

	fmt.Println(w.MediumWordChains(arg1, arg2))
}
