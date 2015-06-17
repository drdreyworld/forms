package forms

func init() {
	FieldsFactory.Register("text", new(TextField))
}

type TextField struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype TextField) Create(meta FieldMeta) Field {

	field := new(TextareaField)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "textarea"

	return field
}

func (field TextField) GetLabel() string {
	return field.Label
}

func (field *TextField) SetLabel(label string) {
	field.Label = label
}

func (field TextField) GetName() string {
	return field.Name
}

func (field *TextField) SetName(name string) {
	field.Name = name
}

func (field TextField) GetValue() interface{} {
	return field.Value
}

func (field *TextField) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field TextField) IsValid(value interface{}) (result bool, err *string) {
	val, ok := value.(string)

	// @TODO ok = validators.Validate(val) && ok

	if ok {
		field.Value = val
	}

	return ok, nil
}
