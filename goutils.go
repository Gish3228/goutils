package goutils

func Contains(a []string, x string) bool {
	for _, el := range a {
		if el == x {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, el := range a {
		if el == x {
			return true
		}
	}
	return false
}
