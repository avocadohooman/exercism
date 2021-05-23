package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Error")
	}
	i := 0
	hammingDistance := 0
	for i < len(a) {
		if a[i] != b[i] {
			hammingDistance += 1
		}
		i++
	}
	return hammingDistance, nil
}
