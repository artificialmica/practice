package practice

func FirstWord(s string) string {
	// Find the start of the first word by skipping leading spaces
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Find the end of the first word
	j := i
	for j < len(s) && s[j] != ' ' {
		j++
	}

	// Extract and return the first word followed by a newline
	return s[i:j] + "\n"
}