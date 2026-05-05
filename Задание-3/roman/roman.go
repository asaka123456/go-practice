package roman

func ToArabic(roman string) int {
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	result := 0
	for i := 0; i < len(roman); i++ {
		current := values[roman[i]]
		if i+1 < len(roman) && values[roman[i+1]] > current {
			result -= current
		} else {
			result += current
		}
	}
	return result
}
