package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func main() {
	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	//hash, err := generateFromPassword("", p)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

// Хеширование пароля с base64-кодированием
func GenerateFromPassword(password string) (hash string, err error) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	// Генерация соли
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	// Хеширование пароля
	hashBytes := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Кодируем результат в base64 для вставки в строковое поле
	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashStr := base64.StdEncoding.EncodeToString(hashBytes)

	// Возвращаем соль и хеш как одну строку (или можно хранить отдельно)
	res := fmt.Sprintf("%s.%s", saltStr, hashStr)

	res = strings.Replace(res, "==.", "", -1)
	//res = strings.Replace(res, "=.", "", -1)
	return res, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
