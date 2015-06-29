package fields

func init() {
	Factory.Register("button", new(Button))
}

type Button struct {
	Name  string
	Label string
	Value string
	Type  string
	Order int
}

func (prototype Button) Create(meta FieldMeta) Field {

	field := new(Button)

	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)

	field.Type = "button"

	return field
}

func (field Button) GetType() string {
	return "button"
}

func (field Button) GetLabel() string {
	return field.Label
}

func (field *Button) SetLabel(label string) {
	field.Label = label
}

func (field Button) GetName() string {
	return field.Name
}

func (field *Button) SetName(name string) {
	field.Name = name
}

func (field Button) GetValue() interface{} {
	return field.Value
}

func (field *Button) SetValue(value interface{}) {
	field.Value = value.(string)
}

func (field Button) IsValid(value interface{}) (result bool, err *string) {
	return true, nil
}

func (field *Button) GetOrder() int {
	return field.Order
}

func (field *Button) SetOrder(order int) {
	field.Order = order
}
