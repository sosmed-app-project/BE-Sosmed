package helper

import gonanoid "github.com/matoous/go-nanoid"

func GenerateId() (string, error) {
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newid, errID := gonanoid.Generate(str, 20)
	if errID != nil {
		return "", errID
	}
	return newid, nil
}
