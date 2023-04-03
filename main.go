package main

import (
	"fmt"
	"strconv"
)

/*
	func arabic(a, b int) int {
		var oper string
		fmt.Println("input")
		fmt.Scan(&a, &oper, &b)
		var res int
		if a < 1 {
			fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от 1 до 10")
		}
		if a > 10 {
			fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от 1 до 10")
		}
		if b < 1 {
			fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от 1 до 10")
		}
		if b > 10 {
			fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от 1 до 10")
		}

		if oper == "+" {
			res = a + b
		}
		if oper == "-" {
			res = a - b
		}
		if oper == "*" {
			res = a * b
		}
		if oper == "/" {
			res = a / b
		}
		return res
	}
*/
func rim(a, b string) (oper string) {
	fmt.Println("input")
	fmt.Scan(&a, &oper, &b)
	_, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("%q looks like a number.\n", a)
	}
	if _, err := strconv.Atoi(a); err == nil {
		fmt.Println("%q looks like a number.\n", b)
	}
	var res string
	var carta = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XX": 20, "XXX": 30, "XL": 40, "L": 50,
		"LX": 60, "LXX": 70, "LXXX": 80, "XC": 90,
		"C": 100}
	var arcarta int
	var arcartab int
	arcarta = carta[a]
	arcartab = carta[b]

	if arcarta < 1 {
		fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от I до X")
	}
	if arcartab > 10 {
		fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от I до X")
	}
	if arcartab < 1 {
		fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от I до X")
	}
	if arcarta > 10 {
		fmt.Println("калькулятор пока что может выполнять операции только с целыми числами от I до X")
	}
	for key := range carta {
		var ress int
		var arcarta int
		var arcartab int
		arcarta = carta[a]
		arcartab = carta[b]
		key = key
		if oper == "+" {
			ress = arcarta + arcartab
		}
		if oper == "-" {
			ress = arcarta - arcartab
		}
		if oper == "*" {
			ress = arcarta * arcartab
		}
		if oper == "/" {
			ress = arcarta / arcartab
		}

		var res int
		var sip int
		res = ress % 10
		sip = ress - res

		var keys string
		var keyss string
		for key, i := range carta {
			if i == sip {
				keys = key
			}
			for key, i := range carta {
				if i == res {
					keyss = key
				}
				result := keys + keyss
				fmt.Println(result)

			}
		}
	}
	return res
}

func main() {
	//value := arabic(1, 1)
	//fmt.Println(value)
	valuer := rim("", "")
	fmt.Println(valuer)
}
