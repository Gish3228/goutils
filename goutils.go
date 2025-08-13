package goutils

func InString(a []string, x string) bool {
	for _, el := range a {
		if el == x {
			return true
		}
	}
	return false
}

func InInt(a []int, x int) bool {
	for _, el := range a {
		if el == x {
			return true
		}
	}
	return false
}
