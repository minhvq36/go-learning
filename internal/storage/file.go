package storage

import "os"

func SaveNote(fileName, content string) error {
	data := []byte(content)

	return os.WriteFile(fileName, data, 0644)
}

func LoadNote(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
