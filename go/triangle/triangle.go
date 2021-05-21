package triangle

import (
	"math"
)

type Kind int

const (
    NaT Kind = iota // 0 = not a triangle
    Equ // 1 = equilateral
    Iso // 2 = isosceles
    Sca // 3 = scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind;
	// First I am checking if the numbers are valid numbers (e.g. inifinte number or just invalid)
	if (!checkValidNumber(a, b, c)) {
		k = NaT;
		return k;
	}
	// Then I check for Triangle Equality
	if (!checkTriangleInequality(a, b, c)) {
		k = NaT;
		return k;
	}
	// Finally I asses the type of triangle
	k = assesTriangle(a, b, c, k);
	return k
}


func checkValidNumber(a, b, c float64) bool {
	if (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)) {
		return false;
	}
	return true;
} 

func checkTriangleInequality(a, b, c float64) bool {
	if (a <= 0 || b <= 0 || c <= 0) {
		return false;
	}
	if ((a + b) < c || (b + c) < a || (a + c) < b) {
		return false;
	}
	return true;
} 

func assesTriangle(a, b, c float64, k Kind) Kind {
	var array [3]float64;
	array[0] = a;
	array[1] = b;
	array[2] = c;
	countEqu := 0;
	countIso := 0;
	for i := 0; i < len(array); i++ {
		same := 0;
		for z := 0; z < len(array); z++ {
			if (array[i] == array[z] && i != z) {
				same++;
			}
		}
		if (same == 2) {
			countEqu++;
		} else if (same == 1) {
			countIso++;
		}
	}
	if (countEqu == 3) {
		k = Equ;
	} else if (countIso == 2) {
		k = Iso;
	} else {
		k = Sca;
	}
	return k;
}
