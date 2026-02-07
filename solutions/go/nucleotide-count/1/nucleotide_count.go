package dna


import "fmt"
// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
    for _, val := range d {
        n, ok := h[val]
        if ok {
            h[val] = n + 1
        } else {
            return nil, fmt.Errorf("INVALID")
        }
    }
	return h, nil
}
