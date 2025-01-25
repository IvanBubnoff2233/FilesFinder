package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

func searchFiles(root string, extension string) ([]string, error) {
	var files []string

	// Функция, которая будет вызываться для каждого файла/папки
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Если произошла ошибка при доступе к файлу/папке, пропускаем
			fmt.Printf("Не удалось получить доступ к %s: %v\n", path, err)
			return nil
		}
		// Проверяем, является ли элемент файлом и имеет ли он нужное расширение
		if !info.IsDir() && strings.EqualFold(filepath.Ext(info.Name()), extension) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func run() bool {
	res, err := askUserRootFile()
	if err != nil {
		fmt.Println(err)
		return true
	}

	direct := "C:\\"

	sliceFiles, err := searchFiles(direct, res)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range sliceFiles {
		fmt.Println(v)
	}

	fmt.Println("Введите e чтобы выйти ")
	var e string
	fmt.Scan(&e)
	if e == "e" {
		return false
	}

	return true
}

func main() {
	for run() {
	}
}
