package fields

import (
	"github.com/drdreyworld/forms/validators"
)

func init() {
	Factory.Register("textarea", new(Textarea))
}

type Textarea struct {
	Name       string
	Label      string
	Value      string
	Type       string
	Order      int
	Error      string
	Validators validators.Validators
}

func (prototype Textarea) Create(meta FieldMeta) Field {
	field := new(Textarea)
	field.Type = "textarea"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *Textarea) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field Textarea) GetError() string {
	return field.Error
}

func (field Textarea) GetType() string {
	return "textarea"
}

func (field Textarea) GetLabel() string {
	return field.Label
}

func (field *Textarea) SetLabel(label string) {
	field.Label = label
}

func (field Textarea) GetName() string {
	return field.Name
}

func (field *Textarea) SetName(name string) {
	field.Name = name
}

func (field Textarea) GetValue() interface{} {
	return field.Value
}

func (field *Textarea) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field *Textarea) IsValid(value interface{}) (result bool, err *string) {
	field.Value, result = value.(string)
	for _, validator := range field.Validators {
		if result = validator.IsValid(value); !result {
			field.Error = validator.GetError()
			break
		}
	}
	return result, nil
}

func (field *Textarea) GetOrder() int {
	return field.Order
}

func (field *Textarea) SetOrder(order int) {
	field.Order = order
}
