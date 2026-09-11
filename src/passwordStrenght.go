package main

func isValidPassword(password string) bool {
	// ?
	
	hasUpper := false
	hasDigit := false
	hasLen := len(password) >= 5 && len(password) <= 12

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	return hasLen && hasUpper && hasDigit
}
