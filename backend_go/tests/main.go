package main

import (
	"fmt"

	"github.com/VladislavSCV/pkg"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			pkg.Log(fmt.Errorf("panic: %v", r))
		}
	}()
	var num1, num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	num, err := Devide(num1, num2)
	if pkg.LogWriteFileReturnError(err) != nil { // Проверяем ошибку после вызова LogWriteFileReturnError
		return
	}
	fmt.Println(num)
	fmt.Println("Hello")
}

func Devide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("division by zero") // Возвращаем ошибку, если делитель ноль
	}
	return num1 / num2, nil
}
