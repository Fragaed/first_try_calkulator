package main

import (
	"fmt"
	"strconv"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func convertToArabic(romanNumeral string) (int, error) {
	number, ok := romanNumerals[romanNumeral]
	if !ok {
		return 0, fmt.Errorf("Invalid roman numeral: %s", romanNumeral)
	}
	return number, nil
}

func convertToRoman(arabicNumeral int) (string, error) {
	if arabicNumeral < 1 || arabicNumeral > 10 {
		return "", fmt.Errorf("Invalid arabic number: %d", arabicNumeral)
	}
	romanNumeral, _ := arabicNumerals[arabicNumeral]
	return romanNumeral, nil
}

func calculate(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}

func main() {
	var aStr, operator, bStr string
	fmt.Println("Enter the expression (for example, 2 + 3): ")
	fmt.Scanln(&aStr, &operator, &bStr)

	a, err := strconv.Atoi(aStr)
	var isRoman bool
	if err != nil {
		isRoman = true
		a, err = convertToArabic(aStr)
		if err != nil {
			panic(err)
		}
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		isRoman = true
		b, err = convertToArabic(bStr)
		if err != nil {
			panic(err)
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Input numbers should be between 1 and 10")
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	if isRoman {
		romanResult, err := convertToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Result:", romanResult)
	} else {
		fmt.Println("Result:", result)
	}
}
