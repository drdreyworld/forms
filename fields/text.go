package fields

func init() {
	Factory.Register("text", new(Text))
}

type Text struct {
	BaseField
	Value string
}

func (prototype Text) Create(meta FieldMeta) Field {
	field := new(Text)
	field.Type = "text"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Text) GetValue() interface{} {
	return field.Value
}

func (field *Text) SetValue(value interface{}) (ok bool) {
	field.Value, ok = value.(string)
	return ok
}

func (field *Text) IsValid(value interface{}) (result bool) {
	if result = field.SetValue(value); result {
		for _, validator := range field.Validators {
			if result = validator.IsValid(value); !result {
				field.Error = validator.GetError()
				break
			}
		}
	}
	return result
}
