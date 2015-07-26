package fields

import (
	"github.com/drdreyworld/forms/validators"
	"time"
)

func init() {
	Factory.Register("datetime", new(DateTime))
}

type DateTime struct {
	Name       string
	Label      string
	Value      string
	Type       string
	Order      int
	Error      string
	Validators validators.Validators
}

func (prototype DateTime) Create(meta FieldMeta) Field {
	field := new(DateTime)
	field.Type = "datetime"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *DateTime) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field DateTime) GetError() string {
	return field.Error
}

func (field DateTime) GetType() string {
	return "datetime"
}

func (field DateTime) GetLabel() string {
	return field.Label
}

func (field *DateTime) SetLabel(label string) {
	field.Label = label
}

func (field DateTime) GetName() string {
	return field.Name
}

func (field *DateTime) SetName(name string) {
	field.Name = name
}

func (field DateTime) GetValue() interface{} {
	return field.Value
}

func (field DateTime) GetValueInHumanFormat() interface{} {
	if len(field.Value) == 0 {
		return ""
	}

	if t, err := time.Parse("2006-01-02 15:04:05", field.Value); err == nil {
		return t.Format("02.01.2006 15:04")
	}
	return field.Value
}

func (field *DateTime) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field *DateTime) IsValid(value interface{}) (result bool, err *string) {
	field.Value, result = value.(string)
	for _, validator := range field.Validators {
		if result = validator.IsValid(value); !result {
			field.Error = validator.GetError()
			break
		}
	}
	if result {
		t, _ := time.Parse("02.01.2006 15:04", field.Value)
		field.Value = t.Format("2006-01-02 15:04:05")
	}
	return result, nil
}

func (field *DateTime) GetOrder() int {
	return field.Order
}

func (field *DateTime) SetOrder(order int) {
	field.Order = order
}
