package main

import "fmt"

func main() {
    // Константы конвертации
    const usdToEur = 0.92   // 1 USD = 0.92 EUR
    const usdToRub = 98.50  // 1 USD = 98.50 RUB
    
    // Курс EUR к RUB (рассчитывается через USD)
    eurToRub := usdToRub / usdToEur
    
    // Пример конвертации
    var usdAmount float64 = 100.0
    
    eurAmount := usdAmount * usdToEur
    rubAmount := usdAmount * usdToRub
    
    // Вывод результатов
    fmt.Println("=== Калькулятор конвертации валют ===")
    fmt.Printf("Курсы:\n")
    fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
    fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
    fmt.Printf("1 EUR = %.2f RUB (рассчитано)\n\n", eurToRub)
    
    fmt.Printf("Конвертация %.2f USD:\n", usdAmount)
    fmt.Printf("-> %.2f EUR\n", eurAmount)
    fmt.Printf("-> %.2f RUB\n", rubAmount)
    
    // Конвертация EUR в RUB
    var eurAmount2 float64 = 50.0
    rubFromEur := eurAmount2 * eurToRub
    fmt.Printf("\n%.2f EUR = %.2f RUB\n", eurAmount2, rubFromEur)
}