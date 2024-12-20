package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	currentWidth := 0

	for _, word := range words {
		if currentWidth+len(word)+len(currentLine) > maxWidth {
			spacesToAdd := maxWidth - currentWidth
			numSpaces := len(currentLine) - 1

			if numSpaces == 0 { // only one word per line
				result = append(result, currentLine[0]+strings.Repeat(" ", spacesToAdd))
			} else {
				spaceBetweenWords := spacesToAdd / numSpaces
				extraSpaces := spacesToAdd % numSpaces
				line := currentLine[0]
				for i := 1; i < len(currentLine); i++ {
					spaces := spaceBetweenWords
					if i <= extraSpaces {
						spaces++
					}

					line += strings.Repeat(" ", spaces) + currentLine[i]
				}
				result = append(result, line)
			}

			// reset before go to next line
			currentLine = nil
			currentWidth = 0
		}

		currentLine = append(currentLine, word)
		currentWidth += len(word)
	}

	lastLine := strings.Join(currentLine, " ") + strings.Repeat(" ", maxWidth-currentWidth-(len(currentLine)-1))
	result = append(result, lastLine)

	return result
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	fmt.Println(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}
