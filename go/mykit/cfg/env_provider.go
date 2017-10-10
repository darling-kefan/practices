package cfg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// EnvProvider provides configuration from the enviroment. All keys will be
// made uppercase.
type EnvProvider struct {
	Namespace string
}

// Provide implements the Provider interface.
func (ep EnvProvider) Provide() (map[string]string, error) {
	config := make(map[string]string)
	
	envs := os.Environ()
	if len(envs) == 0 {
		return nil, errors.New("No enviroment variables found")
	}

	uspace := fmt.Sprintf("%s_", strings.ToUpper(ep.Namespace))
	for _, item := range envs {
		if !strings.HasPrefix(item, uspace) {
			continue
		}
		idx := strings.Index(item, "=")
		config[strings.ToUpper(strings.TrimPrefix(item[:idx], uspace))] = item[idx+1:]
	}

	if len(config) == 0 {
		return nil, fmt.Errorf("Namespace %q was not found", ep.Namespace)
	}

	return config, nil
}
