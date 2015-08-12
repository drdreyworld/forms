package fields

import "github.com/drdreyworld/forms/validators"

type BaseField struct {
	Name       string
	Label      string
	Type       string
	Order      int
	Error      string
	Validators validators.Validators
}

func (prototype BaseField) Create(meta FieldMeta) Field {
	panic("Base field can't be created directly!")
}

func (field *BaseField) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field BaseField) GetError() string {
	return field.Error
}

func (field BaseField) GetLabel() string {
	return field.Label
}

func (field *BaseField) SetLabel(label string) {
	field.Label = label
}

func (field BaseField) GetType() string {
	return field.Type
}

func (field BaseField) GetName() string {
	return field.Name
}

func (field *BaseField) SetName(name string) {
	field.Name = name
}

func (field *BaseField) GetOrder() int {
	return field.Order
}

func (field *BaseField) SetOrder(order int) {
	field.Order = order
}
