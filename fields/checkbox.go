package fields

import (
	"fmt"
	"github.com/drdreyworld/forms/validators"
)

func init() {
	Factory.Register("checkbox", new(Checkbox))
}

type Checkbox struct {
	Name       string
	Label      string
	Value      int
	Type       string
	Order      int
	Error      string
	Validators validators.Validators
}

func (prototype Checkbox) Create(meta FieldMeta) Field {
	field := new(Checkbox)
	field.Type = "checkbox"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *Checkbox) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field Checkbox) GetError() string {
	return field.Error
}

func (field Checkbox) GetType() string {
	return "checkbox"
}

func (field Checkbox) GetLabel() string {
	return field.Label
}

func (field *Checkbox) SetLabel(label string) {
	field.Label = label
}

func (field Checkbox) GetName() string {
	return field.Name
}

func (field *Checkbox) SetName(name string) {
	field.Name = name
}

func (field Checkbox) GetValue() interface{} {
	return field.Value
}

func (field *Checkbox) SetValue(value interface{}) {
	var ok bool
	if field.Value, ok = value.(int); !ok {
		fmt.Println("fuck checkbox")
	}
}

func (field *Checkbox) IsValid(value interface{}) (result bool, err *string) {
	if value.(string) == "1" {
		field.Value = 1
	} else {
		field.Value = 0
	}
	return true, nil
}

func (field *Checkbox) GetOrder() int {
	return field.Order
}

func (field *Checkbox) SetOrder(order int) {
	field.Order = order
}
