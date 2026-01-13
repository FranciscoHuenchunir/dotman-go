package validate

import (
	"fmt"
	"strings"
)

func CleanUniqueArgs(args []string) ([]string, error) {
	var validNames []string
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			return nil, fmt.Errorf("el argumento está vacío o solo tiene espacios")
		}

		validNames = append(validNames, arg)
	}

	return validNames, nil
}
