
package practice

func HashCode(dec string) string {
  
    size := len(dec)
	hashed := make([]byte, len(dec))

	for i, char := range dec {
		hash := (int(char) + size) % 127 // Compute hash
		if hash < 33 {                  // Ensure printable ASCII
			hash += 33
		}
		hashed[i] = byte(hash) // Store the hashed character
	}

	return string(hashed)
}



