package config

import (
	"fmt"
	"os"
	"strings"
)

func GetStrings(config map[string]any, fields ...string) (map[string]string, error) {
	out := make(map[string]string)
	for _, field := range fields {
		name := field
		val, ok := config[field].(string)
		if !ok || val == "" {
			return nil, fmt.Errorf("missing '%s' config field", field)
		}

		if strings.HasSuffix(field, "_envvar") {
			name = strings.TrimSuffix(field, "_envvar")
			envVal := val
			if val, ok = os.LookupEnv(envVal); !ok {
				return nil, fmt.Errorf("missing '%s' environment variable (for '%s' config field)", envVal, field)
			}
		}

		if val == "" {
			return nil, fmt.Errorf("empty '%s' config field", field)
		}

		out[name] = val
	}
	return out, nil
}
