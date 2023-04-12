package main

import (
	"fmt"
	"strconv"
	"strings"
)

// main вызывает функцию input, по сути-обработчик
func input() []string {
	fmt.Println("Input")
	a := ""
	fmt.Scan(&a)
	simvol := strings.Split(a, "")
	// Считаем количество мат операторов, пока что коряво, без функции
	if strings.Count(a, "+") > 1 || strings.Count(a, "-") > 1 || strings.Count(a, "/") > 1 || strings.Count(a, "*") > 1 {
		fmt.Println("ошибка")

	} else if strings.Count(a, "+") == 1 && strings.Count(a, "-") == 1 {
		fmt.Println("ошибка")
	} else if strings.Count(a, "+") == 1 && strings.Count(a, "/") == 1 {
		fmt.Println("ошибка")
	} else if strings.Count(a, "+") == 1 && strings.Count(a, "*") == 1 {
		fmt.Println("ошибка")
	} else if strings.Count(a, "-") == 1 && strings.Count(a, "/") == 1 {
		fmt.Println("ошибка")
	} else if strings.Count(a, "-") == 1 && strings.Count(a, "*") == 1 {
		fmt.Println("ошибка")
	} else if strings.Count(a, "/") == 1 && strings.Count(a, "*") == 1 {
		fmt.Println("ошибка")
	} else {

		//собираем из массива строки
		num_1 := ""
		math := ""
		num_2 := ""
		lenght := len(simvol)
		for k := 0; k < lenght; k++ {
			if simvol[k] == "+" {
				math += simvol[k]
				break
			} else if simvol[k] == "-" {
				math += simvol[k]
				break
			} else if simvol[k] == "/" {
				math += simvol[k]
				break
			} else if simvol[k] == "*" {
				math += simvol[k]
				break

			} else {
				num_1 += simvol[k]
			}
		}

		for k := 0; k < lenght; k++ {
			if simvol[k] == math {
				for i := k + 1; i < lenght; i++ {
					num_2 += simvol[i]
				}
			}
		}

		num1, err1 := strconv.Atoi(num_1)
		num2, err2 := strconv.Atoi(num_2)
		// преобразуем в число
		if err1 == nil && err2 == nil {
			//if len(simvol) > 3 {
			//	fmt.Println("Пожалуйста, используйте два целых числа от 1 до 10")
			//}
			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				fmt.Println("Пожалуйста, используйте два целых числа от 1 до 10")
			} else {
				if math == "+" {
					f := func(err1 int, err2 int) int { return err1 + err2 }
					fmt.Println(f(num1, num2))

				} else if simvol[1] == "-" {
					f := func(err1 int, err2 int) int { return err1 - err2 }
					fmt.Println(f(num1, num2))

				} else if simvol[1] == "*" {
					f := func(err1 int, err2 int) int { return err1 * err2 }
					fmt.Println(f(num1, num2))

				} else if simvol[1] == "/" {
					f := func(err1 int, err2 int) int { return err1 / err2 }
					fmt.Println(f(num1, num2))
				} else {
					fmt.Println("Некорректный математический оператор")
				}
			}
			// создаем карту
		} else if err1 != nil && err2 != nil {
			var carta = map[string]int{
				"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
				"XX": 20, "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100}

			arcart_a := carta[num_1]
			arcarta_b := carta[num_2]

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

				for key := range carta {
					arcart_a = carta[num_1]
					arcarta_b = carta[num_2]
					key = key
					if math == "+" {
						ress = arcart_a + arcarta_b
					} else if math == "-" {
						ress = arcart_a - arcarta_b
					} else if math == "*" {
						ress = arcart_a * arcarta_b
					} else if math == "/" {
						ress = arcart_a / arcarta_b
					} else {
						fmt.Println("Некорректный математический оператор")
					}

				}
				if ress <= 0 {
					fmt.Println("Результат меньше или равен нулю")
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
	}
	return simvol
}
func main() {
	input()

}
