package main

// chop1 run until there is one or no number left
func chop1(numberToFind int, listOfNumbers []int) int {
	length := len(listOfNumbers)
	var begin int

	for begin < length {
		// if number is found or if number is on left side of chop else right side, reset this variables and run until number is find
		if numberToFind == listOfNumbers[(begin+length)/2] {
			return (begin + length) / 2
		} else if numberToFind < listOfNumbers[(begin+length)/2] {
			length = (begin + length) / 2
		} else {
			begin = (begin + length) / 2
		}
	}

	return -1
}
