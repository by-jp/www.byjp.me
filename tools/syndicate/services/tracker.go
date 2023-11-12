package services

import (
	"fmt"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type List struct {
	available map[string]shared.Service
}

func New() *List {
	return &List{}
}

func (l *List) Load(name string, siteConfig any) (shared.Service, error) {
	svc, err := shared.Load(name, siteConfig)
	if err != nil {
		return nil, fmt.Errorf("couldn't load service %s: %w", name, err)
	}
	l.available[name] = svc
	return svc, nil
}

func (l *List) Init(service string) error {
	if svc, ok := l.available[service]; ok {
		return svc.Connect(false)
	}
	return fmt.Errorf("service %s not available", service)
}

func (l *List) Service(name string) shared.Service {
	return l.available[name]
}
