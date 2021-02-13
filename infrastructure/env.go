package infrastructure

import (
	"bufio"
	"os"
	"strings"

	"github.com/justpoypoy/go-base/usecases"
)

// Load is load configs from a env file.
func Load(logger usecases.Logger) {
	path := ".env"

	f, err := os.Open(path)
	if err != nil {
		logger.LogError("%s", err)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.LogError("%s", err)
	}

	for _, l := range lines {
		if l == "" {
			continue
		}
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
}
