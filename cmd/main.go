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

func askDirectory() string {
	fmt.Print("Введите директорию для поиска файла (например C:\\ или C:\\...) ")
	var dir string
	fmt.Scan(&dir)

	return dir
}

func ExitProgram() bool {
	fmt.Println("Вы желаете выйти из программы? y(да) или n(нет)")
	var e string
	fmt.Scan(&e)

	input := strings.ToLower(e)

	if input == "y" {
		return false
	}

	return true
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

		if strings.Contains(path, "Windows") || strings.Contains(path, "Recovery") || strings.Contains(path, "System32") {
			return filepath.SkipDir // Пропускаем эти папки
		}
		return nil
	})
	return files, err
}

func run() bool {
	//Получаем расширение от пользователя и обрабатываем ошибки
	var resUserRoot string
	for {
		res, err := askUserRootFile()
		if err != nil {
			fmt.Println(err)
		} else {
			resUserRoot = res
			break
		}
	}

	//Получаем директорию для поиска файла с расширением пользователя и обратывываем ошибки
	var dir string
	for {
		dir = askDirectory()
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("Данная директория не существует. Попробуйте еще раз...")
		} else {
			break
		}
	}

	//Запускаем поиск файла в системе пользователя
	sliceFiles, err := searchFiles(dir, resUserRoot)
	if err != nil {
		fmt.Println(err)
	}

	//Выводим список путей найденных файлов
	for _, v := range sliceFiles {
		fmt.Println(v)
	}

	//Выход из программы
	ext := ExitProgram()

	return ext
}

func main() {
	fmt.Println(``)
	for run() {
	}
}
