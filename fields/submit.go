package fields

func init() {
	Factory.Register("submit", new(Submit))
}

type Submit struct {
	BaseField
	Value string
}

func (prototype Submit) Create(meta FieldMeta) Field {
	field := new(Submit)
	field.Type = "submit"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Submit) GetValue() interface{} {
	return field.Value
}

func (field *Submit) SetValue(value interface{}) (ok bool) {
	field.Value, ok = value.(string)
	return ok
}

func (field Submit) IsValid(value interface{}) (result bool) {
	return field.SetValue(value)
}
