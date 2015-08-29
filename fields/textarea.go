package fields

func init() {
	Factory.Register("textarea", new(Textarea))
}

type Textarea struct {
	BaseField
	Value string
}

func (prototype Textarea) Create(meta FieldMeta) Field {
	field := new(Textarea)
	field.Type = "textarea"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Textarea) GetValue() interface{} {
	return field.Value
}

func (field *Textarea) SetValue(value interface{}) (ok bool) {
	field.Value, ok = value.(string)
	return ok
}

func (field *Textarea) IsValid(value interface{}) (result bool) {
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
