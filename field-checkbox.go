package forms

func init() {
	FieldsFactory.Register("checkbox", new(CheckboxField))
}

type CheckboxField struct {
	Name  string
	Label string
	Value int
	Type  string
}

func (prototype CheckboxField) Create(meta FieldMeta) Field {

	field := new(CheckboxField)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "checkbox"

	return field
}

func (field CheckboxField) GetType() string {
	return "checkbox"
}

func (field CheckboxField) GetLabel() string {
	return field.Label
}

func (field *CheckboxField) SetLabel(label string) {
	field.Label = label
}

func (field CheckboxField) GetName() string {
	return field.Name
}

func (field *CheckboxField) SetName(name string) {
	field.Name = name
}

func (field CheckboxField) GetValue() interface{} {
	return field.Value
}

func (field *CheckboxField) SetValue(value interface{}) {
	field.Value = value.(int)
}

func (field *CheckboxField) IsValid(value interface{}) (result bool, err *string) {

	if value.(string) == "1" {
		field.Value = 1
	} else {
		field.Value = 0
	}

	return true, nil
}
