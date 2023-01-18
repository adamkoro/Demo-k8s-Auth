package validate

import "regexp"

// Regex to validate email
func Email(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

// Regex to validate username
func Username(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{3,16}$`)
	return re.MatchString(username)
}

// Regex to validate phone number
func Phone(phone string) bool {
	re := regexp.MustCompile(`^[0-9]{10,15}$`)
	return re.MatchString(phone)
}

// Regex to validate name
func Name(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]{3,16}$`)
	return re.MatchString(name)
}
