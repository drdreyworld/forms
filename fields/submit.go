package fields

import (
	"github.com/drdreyworld/forms/validators"
)

func init() {
	Factory.Register("submit", new(Submit))
}

type Submit struct {
	Name  string
	Label string
	Value string
	Type  string
	Order int
}

func (prototype Submit) Create(meta FieldMeta) Field {
	field := new(Submit)
	field.Type = "submit"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *Submit) SetValidators(validators validators.Validators) {}

func (field Submit) GetError() string {
	return ""
}

func (field Submit) GetType() string {
	return "submit"
}

func (field Submit) GetLabel() string {
	return field.Label
}

func (field *Submit) SetLabel(label string) {
	field.Label = label
}

func (field Submit) GetName() string {
	return field.Name
}

func (field *Submit) SetName(name string) {
	field.Name = name
}

func (field Submit) GetValue() interface{} {
	return field.Value
}

func (field *Submit) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field Submit) IsValid(value interface{}) (result bool, err *string) {
	return true, nil
}

func (field *Submit) GetOrder() int {
	return field.Order
}

func (field *Submit) SetOrder(order int) {
	field.Order = order
}
