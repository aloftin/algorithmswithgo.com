package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

// Looping in reverse
// func Reverse(word string) string {
// 	var res string

// 	for i := len(word) - 1; i >= 0; i-- {
// 		res = res + string(word[i])
// 	}

// 	return res
// }

// Looping in reverse with strings.Builder
// func Reverse(word string) string {
// 	var sb strings.Builder

// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}

// 	return sb.String()
// }

// Looping normally with range and runes - handles multi-byte chars in the word
func Reverse(word string) string {
	var res string

	for _, r := range word {
		res = string(r) + res
	}

	return res
}
