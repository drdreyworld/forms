package forms

func init() {
	FieldsFactory.Register("submit", new(SubmitField))
}

type SubmitField struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype SubmitField) Create(meta FieldMeta) Field {

	field := new(SubmitField)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "submit"

	return field
}

func (field SubmitField) GetType() string {
	return "submit"
}

func (field SubmitField) GetLabel() string {
	return field.Label
}

func (field *SubmitField) SetLabel(label string) {
	field.Label = label
}

func (field SubmitField) GetName() string {
	return field.Name
}

func (field *SubmitField) SetName(name string) {
	field.Name = name
}

func (field SubmitField) GetValue() interface{} {
	return field.Value
}

func (field *SubmitField) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field SubmitField) IsValid(value interface{}) (result bool, err *string) {
	return true, nil
}
