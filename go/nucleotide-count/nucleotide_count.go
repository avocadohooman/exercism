package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (d DNA) Counts() (Histogram, error) {
	//first declaring h as Histogram, which is of type map[rune]int
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	//if the len of the DNA is <= 0, then I return an empty map
	if len(d) <= 0 {
		return h, nil
	}
	//looping through the DNA, however, if one of the characters does not match ACGT, I return immediatlley an error
	for i := 0; i < len(d); i++ {
		if d[i] == 'A' {
			h['A'] += 1
		} else if d[i] == 'C' {
			h['C'] += 1
		} else if d[i] == 'G' {
			h['G'] += 1
		} else if d[i] == 'T' {
			h['T'] += 1
		} else {
			return h, errors.New("invalid nucleotides")
		}
	}
	return h, nil
}
