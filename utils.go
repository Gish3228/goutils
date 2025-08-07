package utils

func Contains(a []string, x string) bool {
	for _, el := range a {
		if el == x {
			return true
		}
	}
	return false
}
