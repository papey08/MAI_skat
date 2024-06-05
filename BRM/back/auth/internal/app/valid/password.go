package valid

import "unicode"

func Password(password string) bool {
	if len(password) < 8 || len(password) > 24 {
		return false
	}
	for _, r := range password {
		if !(unicode.Is(unicode.Number, r) || unicode.Is(unicode.Latin, r) || unicode.Is(unicode.Punct, r)) {
			return false
		}
	}
	return true
}
