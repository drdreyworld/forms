package fields

import (
	"github.com/drdreyworld/forms/validators"
)

func init() {
	Factory.Register("text", new(Text))
}

type Text struct {
	Name       string
	Label      string
	Value      string
	Type       string
	Order      int
	Error      string
	Validators validators.Validators
}

func (prototype Text) Create(meta FieldMeta) Field {
	field := new(Text)
	field.Type = "text"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *Text) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field Text) GetError() string {
	return field.Error
}

func (field Text) GetType() string {
	return "text"
}

func (field Text) GetLabel() string {
	return field.Label
}

func (field *Text) SetLabel(label string) {
	field.Label = label
}

func (field Text) GetName() string {
	return field.Name
}

func (field *Text) SetName(name string) {
	field.Name = name
}

func (field Text) GetValue() interface{} {
	return field.Value
}

func (field *Text) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field *Text) IsValid(value interface{}) (result bool, err *string) {
	field.Value, result = value.(string)
	for _, validator := range field.Validators {
		if result = validator.IsValid(value); !result {
			field.Error = validator.GetError()
			break
		}
	}
	return result, nil
}

func (field *Text) GetOrder() int {
	return field.Order
}

func (field *Text) SetOrder(order int) {
	field.Order = order
}
