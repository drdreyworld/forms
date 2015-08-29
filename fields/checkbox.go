package fields

func init() {
	Factory.Register("checkbox", new(Checkbox))
}

type Checkbox struct {
	BaseField
	Value string
}

func (prototype Checkbox) Create(meta FieldMeta) Field {
	field := new(Checkbox)
	field.Type = "checkbox"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Checkbox) GetValue() interface{} {
	return field.Value
}

func (field *Checkbox) SetValue(value interface{}) (ok bool) {
	if field.Value, ok = value.(string); ok {
		if field.Value != "0" {
			field.Value = "1"
		}
	} else {
		field.Value = "0"
	}
	return ok
}

func (field *Checkbox) IsValid(value interface{}) (result bool) {
	return field.SetValue(value)
}
