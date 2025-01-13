package practice


func LastWord(s string) string {
	n := len(s)

	// Start from the end, skipping trailing spaces
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Find the start of the last word
	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}

	// Extract and return the last word followed by a newline
	return s[j+1:i+1] + "\n"
}
