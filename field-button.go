package forms

func init() {
	FieldsFactory.Register("button", new(ButtonField))
}

type ButtonField struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype ButtonField) Create(meta FieldMeta) Field {

	field := new(ButtonField)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "button"

	return field
}

func (field ButtonField) GetType() string {
	return "button"
}

func (field ButtonField) GetLabel() string {
	return field.Label
}

func (field *ButtonField) SetLabel(label string) {
	field.Label = label
}

func (field ButtonField) GetName() string {
	return field.Name
}

func (field *ButtonField) SetName(name string) {
	field.Name = name
}

func (field ButtonField) GetValue() interface{} {
	return field.Value
}

func (field *ButtonField) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field ButtonField) IsValid(value interface{}) (result bool, err *string) {
	return true, nil
}
