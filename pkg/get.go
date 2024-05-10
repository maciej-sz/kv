package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseKeyValueFile(filePath string) (map[string]*Value, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	kv := make(map[string]*Value)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Ignore empty lines or lines starting with #
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := &Value{
			Val:   strings.TrimSpace(parts[1]),
			Quote: "",
		}

		if value.Val[0:1] == "#" {
			continue
		}

		// Remove surrounding quotes if present
		if strings.HasPrefix(value.Val, "\"") && strings.HasSuffix(value.Val, "\"") {
			value.Quote = "\""
		} else if strings.HasPrefix(value.Val, "'") && strings.HasSuffix(value.Val, "'") {
			value.Quote = "'"
		}

		if value.Quote != "" {
			value.Val = value.Val[1 : len(value.Val)-1]
		}

		kv[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return kv, err
}
