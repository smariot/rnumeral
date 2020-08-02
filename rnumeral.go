package rnumeral

import "regexp"

const (
	// Min is the smallest representable roman numeral.
	Min = 1

	// Max is the largest representable roman numeral.
	Max = 3999
)

var (
	re     = regexp.MustCompile(`^(M{0,3})(CD|CM|D?C{0,3})(XL|XC|L?X{0,3})(IV|IX|V?I{0,3})$`)
	lookup = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"XX":   20,
		"XXX":  30,
		"XL":   40,
		"L":    50,
		"LX":   60,
		"LXX":  70,
		"LXXX": 80,
		"XC":   90,
		"C":    100,
		"CC":   200,
		"CCC":  300,
		"CD":   400,
		"D":    500,
		"DC":   600,
		"DCC":  700,
		"DCCC": 800,
		"CM":   900,
		"M":    1000,
		"MM":   2000,
		"MMM":  3000,
	}
)

// Decode parses the given string as a roman numeral, and returns its value.
// Returns 0 for invalid inputs.
func Decode(numeral string) int {
	match := re.FindStringSubmatch(numeral)
	if match == nil {
		return 0
	}

	return lookup[match[1]] + lookup[match[2]] + lookup[match[3]] + lookup[match[4]]
}

var (
	thousands = [4]string{"", "M", "MM", "MMM"}
	hundreds  = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

// Encode returns the given value as roman numeral.
// Returns an empty string if the value is outside the interval [Min, Max].
func Encode(v int) string {
	if v < Min || v > Max {
		return ""
	}

	return thousands[v/1000] + hundreds[v/100%10] + tens[v/10%10] + ones[v%10]
}
