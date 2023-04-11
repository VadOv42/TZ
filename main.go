package main

import (
	"fmt"
	"strconv"
)

// main вызывает функцию input, по сути-обработчик
func input(a, operand, b string) string {
	fmt.Println("Input")
	fmt.Scan(&a, &operand, &b)
	var res string
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)
	// преобразует в число и считает
	if err1 == nil && err2 == nil {
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			fmt.Println("Пожалуйста, используйте два целых числа от 1 до 10")
		} else {
			if operand == "+" {
				f := func(err1 int, err2 int) int { return err1 + err2 }
				fmt.Println(f(num1, num2))

			} else if operand == "-" {
				f := func(err1 int, err2 int) int { return err1 - err2 }
				fmt.Println(f(num1, num2))

			} else if operand == "*" {
				f := func(err1 int, err2 int) int { return err1 * err2 }
				fmt.Println(f(num1, num2))

			} else if operand == "/" {
				f := func(err1 int, err2 int) int { return err1 / err2 }
				fmt.Println(f(num1, num2))
			} else {
				fmt.Println("Некорректный математический оператор")
			}
		}
		// создает карту, ищет значение по ключу
	} else if err1 != nil && err2 != nil {
		var carta = map[string]int{
			"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
			"XX": 20, "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100}
		var arcart_a int
		var arcarta_b int
		arcart_a = carta[a]
		arcarta_b = carta[b]

		if arcart_a < 1 {
			fmt.Println("Пожалуйста, используйте два целых числа от I до X")
		} else if arcarta_b > 10 {
			fmt.Println("Пожалуйста, используйте два целых числа от I до X")
		} else if arcarta_b < 1 {
			fmt.Println("Пожалуйста, используйте два целых числа от I до X")
		} else if arcart_a > 10 {
			fmt.Println("Пожалуйста, используйте два целых числа от I до X")
		} else {
			var ress int
			var arcart_a int
			var arcarta_b int
			fmt.Println(ress)
			// считает
			for key := range carta {
				arcart_a = carta[a]
				arcarta_b = carta[b]
				key = key
				if operand == "+" {
					ress = arcart_a + arcarta_b
				} else if operand == "-" {
					ress = arcart_a - arcarta_b
				} else if operand == "*" {
					ress = arcart_a * arcarta_b
				} else if operand == "/" {
					ress = arcart_a / arcarta_b
				} else {
					fmt.Println("Некорректный математический оператор")
				}

			}
			if ress <= 0 {
				fmt.Println("Результат меньше или равен единице")
			}

			var res int
			var sip int
			res = ress % 10
			sip = ress - res
			var keys string
			var keyss string
			// результат обратно в строку
			for key, i := range carta {
				if i == sip {
					keys = key
				}
			}
			for key, i := range carta {
				if i == res {
					keyss = key
				}

			}

			result := keys + keyss
			fmt.Println(result)

		}
	} else {
		fmt.Println("Нельзя считать арабские с римскими")
	}
	return res
}
func main() {
	input("a", "operand", "b")

}
