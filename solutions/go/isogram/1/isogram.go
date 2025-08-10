package isogram
import "unicode"


func IsIsogram(word string) bool {
    if word == "" {
        return true
    }
	var letterArray [26]bool;
    for _, letter := range word {
        if letter == ' ' || letter == '-' {
            continue
        }
        l := unicode.ToLower(letter)
        if letterArray[l - 97] {
            return false
        } else {
            letterArray[l - 97] = true
        }
    }
    return true
}
