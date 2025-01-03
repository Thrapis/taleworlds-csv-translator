package utils

func CountLeadingSpaces(line string) int {
	count := 0
	for _, v := range line {
		if v == ' ' {
			count++
		} else {
			break
		}
	}

	return count
}

func CountFinalSpaces(line string) int {
	count := 0
	for i := len(line) - 1; i >= 0; i-- {
		v := line[i]

		if v == ' ' {
			count++
		} else {
			break
		}
	}

	return count
}
