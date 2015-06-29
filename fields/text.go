package fields

func init() {
	Factory.Register("text", new(Text))
}

type Text struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype Text) Create(meta FieldMeta) Field {

	field := new(Text)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "text"

	return field
}

func (field Text) GetType() string {
	return "text"
}

func (field Text) GetLabel() string {
	return field.Label
}

func (field *Text) SetLabel(label string) {
	field.Label = label
}

func (field Text) GetName() string {
	return field.Name
}

func (field *Text) SetName(name string) {
	field.Name = name
}

func (field Text) GetValue() interface{} {
	return field.Value
}

func (field *Text) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field *Text) IsValid(value interface{}) (result bool, err *string) {
	val, ok := value.(string)

	// @TODO ok = validators.Validate(val) && ok

	if ok {
		field.Value = val
	}

	return ok, nil
}
