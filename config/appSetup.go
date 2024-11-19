package appSetup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("no s'ha pogut obrir el fitxer: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("línia invàlida al fitxer .env: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error llegint el fitxer: %w", err)
	}

	return nil
}