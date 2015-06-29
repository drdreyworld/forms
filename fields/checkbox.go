package fields

func init() {
	Factory.Register("checkbox", new(Checkbox))
}

type Checkbox struct {
	Name  string
	Label string
	Value int
	Type  string
}

func (prototype Checkbox) Create(meta FieldMeta) Field {

	field := new(Checkbox)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "checkbox"

	return field
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
	field.Value = value.(int)
}

func (field *Checkbox) IsValid(value interface{}) (result bool, err *string) {

	if value.(string) == "1" {
		field.Value = 1
	} else {
		field.Value = 0
	}

	return true, nil
}
