package lab2

func Reverse(s string) string {
	characters := []rune(s)
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
	return string(characters)
}

func PrefixToPostfix(in string) (string, error) {
	return Reverse(in), nil
}
