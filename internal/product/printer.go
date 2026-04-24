package product

import (
	"fmt"
	"strconv"
	"strings"
)

// formatMoney преобразует тийины в строку с пробелами и копейками
func formatMoney(tiins int) string {
	sum := tiins / 100
	rem := tiins % 100

	// Форматируем сумы с пробелами через каждые 3 разряда
	s := strconv.Itoa(sum)
	var parts []string
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		parts = append([]string{s[start:i]}, parts...)
	}
	formattedSum := strings.Join(parts, " ")

	// Добавляем тийины через точку, только если они есть
	if rem > 0 {
		return fmt.Sprintf("%s.%02d", formattedSum, rem)
	}
	return formattedSum
}

// Converter выводит карточку товара. 
// ВАЖНО: сигнатура функции должна быть именно такой для тестов.
func Converter(name, brand string, price int, isAvailable bool, monthly int) {
	fmt.Println("===== Alifshop =====")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %s сум\n", formatMoney(price))
	fmt.Printf("В наличии: %v\n", isAvailable)
	fmt.Printf("Рассрочка: 12 мес → %s сум/мес\n", formatMoney(monthly))
	fmt.Println("====================")
}