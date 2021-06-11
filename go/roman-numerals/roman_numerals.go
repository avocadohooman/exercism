// package romannumerals transforms arabic numbers into roman numbers
package romannumerals

import (
	"errors"
	"strconv"
)

// ArabicOneDigit provides a map for when arabic number is one digit
var ArabicOneDigit = map[byte]string{
	'1': "I",
	'2': "II",
	'3': "III",
	'4': "IV",
	'5': "V",
	'6': "VI",
	'7': "VII",
	'8': "VIII",
	'9': "IX",
}

// ArabicTwoDigits provides a map for when arabic number has two digits
var ArabicTwoDigits = map[byte]string{
	'1': "X",
	'2': "XX",
	'3': "XXX",
	'4': "XL",
	'5': "L",
	'6': "LX",
	'7': "LXX",
	'8': "LXXX",
	'9': "XC",
}

// ArabicThreeDigits provides a map for when arabic number has three digits
var ArabicThreeDigits = map[byte]string{
	'1': "C",
	'2': "CC",
	'3': "CCC",
	'4': "CD",
	'5': "D",
	'6': "DC",
	'7': "DCC",
	'8': "DCCC",
	'9': "CM",
}

// ArabicFourDigits provides a map for when arabic number has four digits
var ArabicFourDigits = map[byte]string{
	'1': "M",
	'2': "MM",
	'3': "MMM",
	'4': "MCD",
	'5': "MD",
	'6': "MDC",
	'7': "MDCC",
	'8': "MDCCC",
	'9': "CM",
}

// ToRomanNumeral takes arabic numbers (0-9) and converts them
// into roman numbers I, V, X, L, C, D, M. Numbers above 3000, or below 0, won't be taken
// into consideration
func ToRomanNumeral(arabic int) (string, error) {
	arabicString := strconv.Itoa(arabic)
	amountOfDigits := NumberOfDigits(arabic)
	romanScore := ""

	if arabic > 3000 || arabic <= 0 {
		return "", errors.New("Error")
	}
	if amountOfDigits == 1 {
		romanScore = ArabicToRomanOneDigit(arabicString)
	} else if amountOfDigits == 2 {
		romanScore = ArabicToRomanTwoDigis(arabicString)
	} else if amountOfDigits == 3 {
		romanScore = ArabicToRomanThreeDigis(arabicString)
	} else {
		romanScore = ArabicToRomanFourDigis(arabicString)
	}
	return romanScore, nil
}

// NumberOfDigits counts the digits of a given number
func NumberOfDigits(arabic int) int {
	digits := 0
	if arabic < 0 {
		arabic *= -1
	}
	if arabic == 0 {
		return 1
	}
	for arabic > 0 {
		arabic /= 10
		digits++
	}
	return digits
}

// ArabicToRomanOneDigit transforms a one digit arabic number
// into roman numbers
func ArabicToRomanOneDigit(arabicString string) string {
	romanScore := ArabicOneDigit[arabicString[0]]
	return romanScore
}

// ArabicToRomanTwoDigis transforms a two digit arabic number
// into roman numbers
func ArabicToRomanTwoDigis(arabicString string) string {
	romanScore := ""
	for i := 0; i < len(arabicString); i++ {
		if i == 0 {
			romanScore += ArabicTwoDigits[arabicString[i]]
		} else {
			romanScore += ArabicOneDigit[arabicString[i]]
		}
	}
	return romanScore
}

// ArabicToRomanThreeDigis transforms a three digit arabic number
// into roman numbers
func ArabicToRomanThreeDigis(arabicString string) string {
	romanScore := ""
	for i := 0; i < len(arabicString); i++ {
		if i == 0 {
			romanScore += ArabicThreeDigits[arabicString[i]]
		} else if i == 1 {
			romanScore += ArabicTwoDigits[arabicString[i]]
		} else {
			romanScore += ArabicOneDigit[arabicString[i]]
		}
	}
	return romanScore
}

// ArabicToRomanFourDigis transforms a four digit arabic number
// into roman numbers
func ArabicToRomanFourDigis(arabicString string) string {
	romanScore := ""
	for i := 0; i < len(arabicString); i++ {
		if i == 0 {
			romanScore += ArabicFourDigits[arabicString[i]]
		} else if i == 1 {
			romanScore += ArabicThreeDigits[arabicString[i]]
		} else if i == 2 {
			romanScore += ArabicTwoDigits[arabicString[i]]
		} else {
			romanScore += ArabicOneDigit[arabicString[i]]
		}
	}
	return romanScore
}
