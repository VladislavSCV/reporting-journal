package pkg

import (
	"fmt"
	"log"
	"os"
)

func LogWriteFileReturnError(err error) error {
	if err != nil {
		Log(err)
		writeErrorInFile(err)
	}
	return err
}

// writeErrorInFile - функция, которая записывает ошибку в файл errors.log.
// Функция создает новый файл, если его не существует, и закрывает его после записи.
// Если не удалось открыть файл, функция выходит, не записывая ошибку.
// Функция не возвращает ошибку и не имеет возвращаемого значения.
func writeErrorInFile(err error) {
	// Открываем файл для записи. Если файла нет, создаем новый.
	file, openErr := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if openErr != nil {
		log.Printf("Не удалось открыть файл для записи ошибок: %v", openErr)
		return // Выходим из функции, если не удалось открыть файл
	}
	defer file.Close() // Закрываем файл после записи

	// Создаем новый логгер, который будет записывать в файл
	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Записываем ошибку в файл
	logger.Println(err)
	fmt.Println("All")
}
