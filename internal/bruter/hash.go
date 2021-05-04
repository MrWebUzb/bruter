package bruter

import (
	"bufio"
	"os"
	"strings"
)

type HashFile struct {
	Username     string
	PasswordHash string
}

func ReadHashFile(filename string) ([]HashFile, error) {
	file, err := os.Open(filename)

	if err != nil {
		return []HashFile{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scan := bufio.NewScanner(reader)

	result := []HashFile{}
	for scan.Scan() {
		text := scan.Text()

		loginWithPassword := strings.Split(text, ":")

		size := len(loginWithPassword)
		if size == 2 {
			result = append(result, HashFile{
				Username:     loginWithPassword[0],
				PasswordHash: loginWithPassword[1],
			})
			continue
		}

		if size == 1 {
			result = append(result, HashFile{
				PasswordHash: loginWithPassword[0],
			})
		}
	}

	return result, nil
}
