package main

import "fmt"

func romanToArabic(roman string) int {
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

func main() {
	fmt.Println("I =", romanToArabic("I"))
	fmt.Println("IV =", romanToArabic("IV"))
	fmt.Println("IX =", romanToArabic("IX"))
	fmt.Println("XL =", romanToArabic("XL"))
	fmt.Println("XC =", romanToArabic("XC"))
	fmt.Println("CD =", romanToArabic("CD"))
	fmt.Println("CM =", romanToArabic("CM"))
	fmt.Println("MCMXCIV =", romanToArabic("MCMXCIV"))
	fmt.Println("MMXXIV =", romanToArabic("MMXXIV"))
}
