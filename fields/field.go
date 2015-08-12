package fields

import (
	"github.com/drdreyworld/forms/validators"
)

type Field interface {
	Create(meta FieldMeta) Field
	GetType() string
	GetLabel() string
	SetLabel(label string)
	GetName() string
	SetName(name string)
	GetValue() interface{}
	SetValue(value interface{}) (ok bool)
	IsValid(value interface{}) (result bool)
	GetOrder() int
	SetOrder(order int)
	SetValidators(validators validators.Validators)
	GetError() string
}

type Fields []Field

func (fields Fields) Len() int {
	return len(fields)
}

func (fields Fields) Less(i, j int) bool {
	return fields[i].GetOrder() < fields[j].GetOrder()
}

func (fields Fields) Swap(i, j int) {
	fields[i], fields[j] = fields[j], fields[i]
}

func (fields Fields) GetField(name string) (Field, bool) {
	for _, field := range fields {
		if field.GetName() == name {
			return field, true
		}
	}
	return nil, false
}

type FieldMeta struct {
	Name       string
	Type       string
	Value      interface{}
	Label      string
	Order      int
	Validators validators.ValidatorsMeta
}

type FieldsMeta []FieldMeta

func (fields FieldsMeta) Len() int {
	return len(fields)
}

func (fields FieldsMeta) Less(i, j int) bool {
	return fields[i].Order < fields[j].Order
}

func (fields FieldsMeta) Swap(i, j int) {
	fields[i], fields[j] = fields[j], fields[i]
}

type ValueOption struct {
	Value interface{}
	Title interface{}
	Order int
}

func (option ValueOption) IsSelected(compare ValueOption) bool {
	return compare.Value == option.Value
}

type ValueOptions []ValueOption

func (fields ValueOptions) Len() int {
	return len(fields)
}

func (fields ValueOptions) Less(i, j int) bool {
	return fields[i].Order < fields[j].Order
}

func (fields ValueOptions) Swap(i, j int) {
	fields[i], fields[j] = fields[j], fields[i]
}

func (options ValueOptions) GetOptionByValue(value interface{}) (ValueOption, bool) {
	for _, option := range options {
		if option.Value == value {
			return option, true
		}
	}
	return ValueOption{}, false
}

type FieldWithOptions interface {
	SetValueOptions(options ValueOptions)
	GetValueOptions() ValueOptions
}
