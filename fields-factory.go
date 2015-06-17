package forms

import (
	"errors"
	"fmt"
)

var FieldsFactory = FieldsFactoryStruct{}

type FieldsFactoryStruct struct {
	registry map[string]Field
}

func (factory *FieldsFactoryStruct) Register(key string, prototype Field) {
	if factory.registry == nil {
		factory.registry = make(map[string]Field)
	}
	factory.registry[key] = prototype
}

func (factory *FieldsFactoryStruct) Unregister(key string) {
	delete(factory.registry, key)
}

func (factory FieldsFactoryStruct) CreateField(meta FieldMeta) (result *Field, err error) {

	prototype, exists := factory.registry[meta.Type]

	if !exists {
		return nil, errors.New(fmt.Sprintf("Field type not registered '%s'", meta.Type))
	}

	field := prototype.Create(meta)

	return &field, nil
}
