package validators

import (
	"errors"
	"fmt"
)

var Factory = ValidatorsFactory{}

type ValidatorsFactory struct {
	registry map[string]Validator
}

func (factory *ValidatorsFactory) Register(key string, prototype Validator) {
	if factory.registry == nil {
		factory.registry = make(map[string]Validator)
	}
	factory.registry[key] = prototype
}

func (factory *ValidatorsFactory) Unregister(key string) {
	delete(factory.registry, key)
}

func (factory ValidatorsFactory) CreateValidator(meta ValidatorMeta) (result Validator, err error) {
	if prototype, exists := factory.registry[meta.Type]; exists {
		return prototype.Create(meta), nil
	}
	return nil, errors.New(fmt.Sprintf("Validator type not registered '%s'", meta.Type, meta))
}
