package bruter

import (
	"bufio"
	"os"
	"strings"
)

func ReadWordlistFile(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scan := bufio.NewScanner(reader)

	result := []string{}
	for scan.Scan() {
		text := scan.Text()
		text = strings.Trim(text, "\n \t")
		result = append(result, text)
	}

	return result, nil
}
