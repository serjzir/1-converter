package main

import (
	"fmt"
	"strings"
)

// Константы конвертации валют
const usdToEur = 0.92  // 1 USD = 0.92 EUR
const usdToRub = 98.50 // 1 USD = 98.50 RUB

// Функция для считывания числа с проверкой
func readAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму для конвертации (положительное число): ")
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка: нужно ввести число. Попробуйте снова.")
			var discard string
			fmt.Scan(&discard) // очищаем буфер
			continue
		}

		if amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть положительной. Попробуйте снова.")
			continue
		}

		return amount
	}
}

// Функция для считывания и проверки валюты
func readCurrency(prompt string) string {
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"RUB": true,
	}

	for {
		fmt.Print(prompt)
		var currency string
		fmt.Scan(&currency)

		// Приводим к верхнему регистру
		currency = strings.ToUpper(currency)

		if validCurrencies[currency] {
			return currency
		}

		fmt.Println("Ошибка: допустимые валюты - USD, EUR, RUB. Попробуйте снова.")
	}
}

// Функция для конвертации валют
func convertCurrency(amount float64, fromCurrency, toCurrency string) float64 {
	// Сначала конвертируем всё в USD (базовая валюта)
	var amountInUSD float64

	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur // EUR → USD
	case "RUB":
		amountInUSD = amount / usdToRub // RUB → USD
	default:
		fmt.Println("Неизвестная валюта источника")
		return 0
	}

	// Конвертируем из USD в целевую валюту
	switch toCurrency {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD * usdToEur
	case "RUB":
		return amountInUSD * usdToRub
	default:
		fmt.Println("Неизвестная целевая валюта")
		return 0
	}
}

// Функция для отображения меню
func showMenu() {
	fmt.Println("\n=== Калькулятор конвертации валют ===")
	fmt.Println("Доступные валюты: USD, EUR, RUB")
	fmt.Println("Курсы:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB (рассчитано)\n\n", usdToRub/usdToEur)
}

// Функция для выполнения одной конвертации
func performConversion() {
	// Шаг 1: Ввод исходной валюты
	fromCurrency := readCurrency("Введите исходную валюту (USD/EUR/RUB): ")

	// Шаг 2: Ввод суммы
	amount := readAmount()

	// Шаг 3: Ввод целевой валюты
	toCurrency := readCurrency("Введите целевую валюту (USD/EUR/RUB): ")

	// Шаг 4: Выполнение конвертации
	result := convertCurrency(amount, fromCurrency, toCurrency)

	// Шаг 5: Вывод результата
	fmt.Printf("\nРезультат конвертации:\n")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func main() {
	for {
		showMenu()

		performConversion()

		// Спрашиваем, хочет ли пользователь продолжить
		fmt.Print("\nХотите выполнить еще одну конвертацию? (да/нет): ")
		var answer string
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)

		if answer != "да" && answer != "yes" && answer != "д" {
			fmt.Println("Спасибо за использование калькулятора валют! До свидания!")
			break
		}
	}
}