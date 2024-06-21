package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Карта для преобразования Римских чисел в Арабские
var romanNumerals = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

// Срез структур для преобразования Арабских чисел в Римские
var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Преобразование римской цифры в целое число
func romanToInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanNumerals[string(s[i])] < romanNumerals[string(s[i+1])] {
			n -= romanNumerals[string(s[i])]
		} else {
			n += romanNumerals[string(s[i])]
		}
	}
	return n
}

// Преобразование целого числа в римскую цифру
func intToRoman(num int) string {
	var result strings.Builder
	for _, r := range arabicToRoman {
		for num >= r.Value {
			result.WriteString(r.Symbol)
			num -= r.Value
		}
	}
	return result.String()
}

// Проверка - является ли строка допустимой римской цифрой
func isRoman(s string) bool {
	for _, r := range s {
		if _, exists := romanNumerals[string(r)]; !exists {
			return false
		}
	}
	return isValidRoman(s)
}

// Проверка на корректность римских цифр от 1 до 10
func isValidRoman(s string) bool {
	validRomans := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true,
		"VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
	}
	return validRomans[s]
}

func main() {
	// Получаем от пользователя входные данные
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 7 + 5 или X * VI):")
	scanner.Scan()
	userInput := scanner.Text()

	// Парсим - разбиваем пользвательский ввод на состовные части
	parts := strings.Fields(userInput)

	// Проверяем на корректный ввод
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	// Определяем - используются при вводе римские или арабские цифры
	isRomanInput := isRoman(aStr) && isRoman(bStr)
	isArabicInput := !isRoman(aStr) && !isRoman(bStr)

	if !isRomanInput && !isArabicInput {
		panic("Смешивание римских и арабских цифр")
	}

	var a, b, result int

	if isRomanInput {
		a = romanToInt(aStr)
		b = romanToInt(bStr)
	} else {
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic("Неверный формат числа")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic("Неверный формат числа")
		}
	}
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10 включительно")
	}

	// Вычисляем выражение
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		result = a / b
	default:
		panic("Неверная операция")
	}

	// Проверка результата для римских цифр
	if isRomanInput && result < 1 {
		panic("Результат меньше единицы для римских чисел")
	}

	// Вывод
	if isRomanInput {
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}
