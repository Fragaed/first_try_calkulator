package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
}

// Функция для конвертации римского числа в арабское
func convertToArabic(romanNumeral string) (int, error) {
	total := 0
	for i := 0; i < len(romanNumeral); {
		if i+1 < len(romanNumeral) && romanNumerals[romanNumeral[i:i+2]] > 0 {
			total += romanNumerals[romanNumeral[i:i+2]]
			i += 2
		} else {
			if romanNumerals[romanNumeral[i:i+1]] > 0 {
				total += romanNumerals[romanNumeral[i:i+1]]
				i++
			} else {
				return 0, fmt.Errorf("invalid Roman numeral: %s", romanNumeral)
			}
		}
	}
	return total, nil
}

// Функция для конвертации арабского числа в римское
func convertToRoman(arabicNumeral int) (string, error) {
	var roman strings.Builder
	arabicToRoman := []struct {
		arabic int
		roman  string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	if arabicNumeral > 100 || arabicNumeral < 1 {
		return "", fmt.Errorf("number out of range for Roman numerals: %d", arabicNumeral)
	}

	for _, pair := range arabicToRoman {
		for arabicNumeral >= pair.arabic {
			roman.WriteString(pair.roman)
			arabicNumeral -= pair.arabic
		}
	}

	return roman.String(), nil
}

// Функция для вычисления результата операции
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
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the expression (for example, II + III or 4 * 5): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read input")
	}

	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Invalid input. Please use the format: <number> <operator> <number>")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	// Инициализация переменных
	var a, b int
	var aIsRoman, bIsRoman bool

	// Проверка и преобразование первого числа
	if a, err = strconv.Atoi(aStr); err == nil {
		aIsRoman = false
	} else if a, err = convertToArabic(aStr); err == nil {
		aIsRoman = true
	} else {
		panic(fmt.Sprintf("Invalid number or Roman numeral: %s", aStr))
	}

	// Проверка и преобразование второго числа
	if b, err = strconv.Atoi(bStr); err == nil {
		bIsRoman = false
	} else if b, err = convertToArabic(bStr); err == nil {
		bIsRoman = true
	} else {
		panic(fmt.Sprintf("Invalid number or Roman numeral: %s", bStr))
	}

	// Проверка, что оба числа либо римские, либо арабские
	if aIsRoman != bIsRoman {
		panic("Cannot perform operations between Roman and Arabic numerals.")
	}

	// Проверка диапазона чисел
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Numbers must be in range from 1 to 10")
	}

	// Вычисление результата
	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	// Вывод результата и конвертация в римское число, если нужно
	if aIsRoman {
		romanResult, err := convertToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Result:", romanResult)
	} else {
		fmt.Println("Result:", result)
	}
}
