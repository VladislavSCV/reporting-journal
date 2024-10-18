package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// CError checks if the error is not nil and logs it. This is a convenience
// function for cases where you want to log an error but don't want to
// explicitly check if the error is not nil.
func CError(err error) error {
	if err != nil {
		Log(err)
		writeErrorInFile(err)
	}
	return err
}

// writeErrorInFile записывает ошибку в файл.
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

func HandleHTTPError(c *gin.Context, statusCode int, message string, err error) {
	// Здесь можно добавить логику логирования ошибок, если это нужно
	err = CError(err)
	if err != nil {
		return
	} // Обрабатываем ошибку, например, логируем её
	c.JSON(statusCode, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}
