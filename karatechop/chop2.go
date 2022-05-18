package main

// doesnt work recursively, the only solution is to create a  global variable ?
func chop2(numberToFind int, listOfNumbers []int) int {
	length := len(listOfNumbers)

	if numberToFind == listOfNumbers[0] {
		return 000000000000
	}

	head := listOfNumbers[0 : length/2]
	tail := listOfNumbers[length/2 : length]

	if numberToFind >= tail[0] {
		return chop2(numberToFind, tail)
	} else if numberToFind <= head[len(head)-1] {
		return chop2(numberToFind, head)
	}

	return -1
}
