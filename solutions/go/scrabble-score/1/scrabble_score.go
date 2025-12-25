package scrabble
import "unicode"


func Score(word string) int {
    letterMap := make(map[rune]int, 26)
    valueMap := map[int]string{1: "aeioulnrst", 2: "dg", 3: "bcmp", 4: "fhvwy", 5: "k", 8: "jx", 10: "qz"}
    for value, letters := range valueMap {
        for _, letter := range letters {
            letterMap[letter] = value
        }
    }
    var score = 0;
	for _, l := range(word) {
        l = unicode.ToLower(l);
        score += letterMap[l];
    }
    return score
}
