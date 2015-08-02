package fields

import (
	"github.com/drdreyworld/forms/validators"
)

func init() {
	Factory.Register("select", new(Select))
}

type Select struct {
	Name         string
	Label        string
	Value        ValueOption
	Type         string
	Order        int
	Error        string
	Validators   validators.Validators
	ValueOptions ValueOptions
}

func (prototype Select) Create(meta FieldMeta) Field {
	field := new(Select)
	field.Type = "select"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field *Select) SetValidators(validators validators.Validators) {
	field.Validators = validators
}

func (field Select) GetError() string {
	return field.Error
}

func (field Select) GetType() string {
	return "select"
}

func (field Select) GetLabel() string {
	return field.Label
}

func (field *Select) SetLabel(label string) {
	field.Label = label
}

func (field Select) GetName() string {
	return field.Name
}

func (field *Select) SetName(name string) {
	field.Name = name
}

func (field Select) GetValue() interface{} {
	return field.Value
}

func (field *Select) SetValue(value interface{}) {
	var ok bool
	if field.Value, ok = field.ValueOptions.GetOptionByValue(value); !ok {
		field.Value = ValueOption{}
	}
}

func (field *Select) IsValid(value interface{}) (result bool, err *string) {
	if len(field.ValueOptions) < 1 {
		return true, nil
	}

	if field.Value, result = field.ValueOptions.GetOptionByValue(value); result {
		for _, validator := range field.Validators {
			if result = validator.IsValid(value); !result {
				field.Error = validator.GetError()
				break
			}
		}
	} else {
		field.Error = "Не допустимый вариант"
	}
	return result, nil
}

func (field *Select) GetOrder() int {
	return field.Order
}

func (field *Select) SetOrder(order int) {
	field.Order = order
}

func (field *Select) SetValueOptions(options ValueOptions) {
	field.ValueOptions = options
}

func (field Select) GetValueOptions() ValueOptions {
	return field.ValueOptions
}
