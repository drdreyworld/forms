package forms

func init() {
	FieldsFactory.Register("textarea", new(TextareaField))
}

type TextareaField struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype TextareaField) Create(meta FieldMeta) Field {

	field := new(TextareaField)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "textarea"

	return field
}

func (field TextareaField) GetLabel() string {
	return field.Label
}

func (field *TextareaField) SetLabel(label string) {
	field.Label = label
}

func (field TextareaField) GetName() string {
	return field.Name
}

func (field *TextareaField) SetName(name string) {
	field.Name = name
}

func (field TextareaField) GetValue() interface{} {
	return field.Value
}

func (field *TextareaField) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field TextareaField) IsValid(value interface{}) (result bool, err *string) {
	val, ok := value.(string)

	// @TODO ok = validators.Validate(val) && ok

	if ok {
		field.Value = val
	}

	return ok, nil
}
