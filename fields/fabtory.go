package fields

import (
	"errors"
	"fmt"
)

var Factory = FieldsFactory{}

type FieldsFactory struct {
	registry map[string]Field
}

func (factory *FieldsFactory) Register(key string, prototype Field) {
	if factory.registry == nil {
		factory.registry = make(map[string]Field)
	}
	factory.registry[key] = prototype
}

func (factory *FieldsFactory) Unregister(key string) {
	delete(factory.registry, key)
}

func (factory FieldsFactory) CreateField(meta FieldMeta) (result *Field, err error) {

	prototype, exists := factory.registry[meta.Type]

	if !exists {
		return nil, errors.New(fmt.Sprintf("Field type not registered '%s'", meta.Type))
	}

	field := prototype.Create(meta)

	return &field, nil
}
