package strand

/*
Given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement:

G -> C
C -> G
T -> A
A -> U
*/

func ToRNA(dna string) string {
	rna := ""

	if len(dna) <= 0 {
		return rna
	}
	for i := 0; i < len(dna); i++ {
		if dna[i] == 'G' {
			rna += "C"
		}
		if dna[i] == 'C' {
			rna += "G"
		}
		if dna[i] == 'T' {
			rna += "A"
		}
		if dna[i] == 'A' {
			rna += "U"
		}
	}
	return rna
}
