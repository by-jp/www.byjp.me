package shared

import (
	"fmt"
)

var serviceTypes = make(map[string]ServiceCreator)

func Register(name string, svc ServiceCreator) {
	serviceTypes[name] = svc
}

func Load(name string, config any) (Service, error) {
	cfgMap, ok := config.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid config for '%s' service", name)
	}

	var serviceType string
	if t, ok := cfgMap["type"].(string); ok {
		serviceType = t
	} else {
		serviceType = name
	}

	svc, ok := serviceTypes[serviceType]
	if !ok {
		return nil, fmt.Errorf("unknown service type: %s", serviceType)
	}

	instance, err := svc(cfgMap)
	if err != nil {
		return nil, fmt.Errorf("unable to create '%s' service: %w", name, err)
	}

	return instance, nil
}
