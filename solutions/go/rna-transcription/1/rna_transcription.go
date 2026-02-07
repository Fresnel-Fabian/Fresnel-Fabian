package strand

import "strings"

func ToRNA(dna string) string {
    var sb strings.Builder
	mapping := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
    for _, val := range dna {
        sb.WriteRune(mapping[val])
    }
    rna := sb.String();
    return rna
}
