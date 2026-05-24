package main

import "fmt"

// Константы конвертации валют
const usdToEur = 0.92  // 1 USD = 0.92 EUR
const usdToRub = 98.50 // 1 USD = 98.50 RUB

// Функция для считывания ввода пользователя
func readAmount() float64 {
	var amount float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	return amount
}

// Функция для конвертации валют
func convertCurrency(amount float64, fromCurrency, toCurrency string) float64 {
	return 0
}

func main() {
	// Курс EUR к RUB (рассчитывается через USD)
	eurToRub := usdToRub / usdToEur

	// Вывод курсов
	fmt.Println("=== Калькулятор конвертации валют ===")
	fmt.Printf("Курсы:\n")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB (рассчитано)\n\n", eurToRub)

	// Считывание ввода от пользователя
	amount := readAmount()

	// Вызов пустой функции (пока возвращает 0)
	result := convertCurrency(amount, "USD", "EUR")
	fmt.Printf("%.2f USD = %.2f EUR (результат пустой функции)\n", amount, result)

	// Конвертация USD → RUB через пустую функцию
	resultRub := convertCurrency(amount, "USD", "RUB")
	fmt.Printf("%.2f USD = %.2f RUB (результат пустой функции)\n", amount, resultRub)
}
