package fields

func init() {
	Factory.Register("textarea", new(Textarea))
}

type Textarea struct {
	Name  string
	Label string
	Value string
	Type  string
}

func (prototype Textarea) Create(meta FieldMeta) Field {

	field := new(Textarea)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "textarea"

	return field
}

func (field Textarea) GetType() string {
	return "textarea"
}

func (field Textarea) GetLabel() string {
	return field.Label
}

func (field *Textarea) SetLabel(label string) {
	field.Label = label
}

func (field Textarea) GetName() string {
	return field.Name
}

func (field *Textarea) SetName(name string) {
	field.Name = name
}

func (field Textarea) GetValue() interface{} {
	return field.Value
}

func (field *Textarea) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field *Textarea) IsValid(value interface{}) (result bool, err *string) {
	val, ok := value.(string)

	// @TODO ok = validators.Validate(val) && ok

	if ok {
		field.Value = val
	}

	return ok, nil
}
