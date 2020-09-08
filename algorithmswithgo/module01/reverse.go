package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var rev string
	for _, val := range word {
		rev = string(val) + rev
	}
	return rev
}

// using runes and string cast
