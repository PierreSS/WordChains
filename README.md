# Description

Write a program that solves word-chain puzzles.

http://codekata.com/kata/kata19-word-chains/

Write a program that do a binary search on list of integers.

http://codekata.com/kata/kata02-karate-chop/

# Getting started

To compile you must have at least go 1.17

Go to the 2 different files and do : make, once it is done, enter ./bin/$exercise

# Explanation

For the first exercise, I used the iterative way for a binary chop, I did a chop and set the indexes on start or mid depending of where the number to find was. I tried to do a second one with a recursion calling back my function shrinking my list until it was equal to 1 number or 0. Unforunatly, I dont think it can work since we can't recall of the global length except if we put a global variable.

For the second exercise, I did a simple word chains to begin with, I read throught the english word file, iterate on all the characters of the final word and copy them into the first one, if the first word is an english word with his letter changed from final word, the word is saved and printed, then iterate until final word is equal to the generated word. For the medium one, I used the same thing but I added a condition if the iteration lead to nowhere, then it restart ignoring the blocking word when checking for a new word to generate. For the final word chains I created a ticker to create a new word with one different letter from the first word, then iterate as medium word chains.