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

func (factory FieldsFactory) CreateField(meta FieldMeta) (result Field, err error) {
	if prototype, exists := factory.registry[meta.Type]; exists {
		return prototype.Create(meta), nil
	}
	return nil, errors.New(fmt.Sprintf("Field type not registered '%s'", meta.Type))
}
