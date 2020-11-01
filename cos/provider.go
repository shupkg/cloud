package cos

import (
	"errors"
	"fmt"
)

var (
	ErrProviderNotFound = errors.New("provider not found")

	providers = map[string]FactoryFunc{}
)

type FactoryFunc func(options Options) Service

func Register(provider string, factory FactoryFunc, override bool) bool {
	if _, find := providers[provider]; find {
		if !override {
			return false
		}
	}
	providers[provider] = factory
	return true
}

func New(provider string, options Options) (Service, error) {
	factory, find := providers[provider]
	if !find {
		return nil, fmt.Errorf("%w: %s", ErrProviderNotFound, provider)
	}
	return factory(options), nil
}
