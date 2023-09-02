// Calc project main.go
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	
)
	var roman = []string{"I", "II", "II", "IV", "V", "VI", "VII", "VIII", "IX", "X", "+", "-", "*", "/"}
	var arabic = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "+", "-", "*", "/"}
	//var digit1 int
	//var digit2 int

func main() {
	var line string
	fmt.Println("Введите два числа 0т 1 до 10 в формате 1+2 либо I*II:")
	fmt.Fscan(os.Stdin, &line)
	line = strings.TrimSpace(line)
	var symbols []string = strings.Split(line, "")
	if len(line) > 2 {
		fmt.Print(errors.New("Введено некорректное количество символов"))
	} else if 

	//fmt.Printf("%q\n", strings.Split(line, ""))
	fmt.Println(symbols)
}
