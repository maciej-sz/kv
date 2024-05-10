package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func SaveToFile(filePath string, data map[string]*Value) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create or open file: %w", err)
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err = writer.Flush()
	}(writer)

	for key, value := range data {
		line := fmt.Sprintf("%s=%s%s%s\n", key, value.Quote, value.Val, value.Quote)
		if _, err := writer.WriteString(line); err != nil {
			return fmt.Errorf("could not write to file: %w", err)
		}
	}

	return err
}
