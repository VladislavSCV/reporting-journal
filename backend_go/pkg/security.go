package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type HashResult struct {
	Salt string
	Hash string
}

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p params

// GenerateFromPassword Хеширование пароля с base64-кодированием
func GenerateHashFromPassword(password string) (*HashResult, error) {
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
		return nil, err
	}

	// Хеширование пароля
	hashBytes := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Кодируем результат в base64 для вставки в строковое поле
	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashStr := base64.StdEncoding.EncodeToString(hashBytes)

	// Возвращаем соль и хеш как одну строку (или можно хранить отдельно)
	//res := fmt.Sprintf("%s.%s", saltStr, hashStr)
	//
	//res = strings.Replace(res, "==.", "", -1)
	//res = strings.Replace(res, "=.", "", -1)
	hr := HashResult{
		Salt: saltStr,
		Hash: hashStr,
	}
	return &hr, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// VerifyPassword проверяет, соответствует ли введенный пароль хешу
func VerifyPassword(password, saltStr, hashStr string) (bool, error) {
	salt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		return false, err
	}

	hashBytes, err := base64.StdEncoding.DecodeString(hashStr)
	if err != nil {
		return false, err
	}

	// Хешируем введенный пароль с той же солью
	newHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return subtle.ConstantTimeCompare(hashBytes, newHash) == 1, nil
}
