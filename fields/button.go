package fields

import "github.com/drdreyworld/forms/validators"

func init() {
	Factory.Register("button", new(Button))
}

type Button struct {
	BaseField
	Value string
}

func (prototype Button) Create(meta FieldMeta) Field {
	field := new(Button)
	field.Type = "button"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Button) SetValidators(validators validators.Validators) {}

func (field Button) GetValue() interface{} {
	return field.Value
}

func (field *Button) SetValue(value interface{}) (ok bool) {
	field.Value, ok = value.(string)
	return ok
}

func (field Button) IsValid(value interface{}) (result bool) {
	return field.SetValue(value)
}
