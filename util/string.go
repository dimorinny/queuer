package util

func IsAllStringsNotEmpty(values ...string) bool {
	for _, val := range values {
		if val == "" {
			return false
		}
	}

	return true
}
