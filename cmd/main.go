package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ERROR_INCORRECT_INPUT = errors.New("Не корретный ввод расширения файла пользователем")
)

func askUserRootFile() (string, error) {
	fmt.Print("Введите расширение файла (например .txt): ")
	var exp string
	fmt.Scan(&exp)

	if !strings.HasPrefix(exp, ".") {
		return fmt.Sprintln("Введите корретно расширение файла как указанно в примере (например .txt)"), ERROR_INCORRECT_INPUT
	}

	return exp, nil
}

func main() {

}
