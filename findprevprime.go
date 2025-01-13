package practice


func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if IsPrime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}

// Function to check if a number is prime
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}