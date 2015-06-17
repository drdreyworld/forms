package forms

type Field interface {
	Create(meta FieldMeta) Field
	GetLabel() string
	SetLabel(label string)
	GetName() string
	SetName(name string)
	GetValue() interface{}
	SetValue(value interface{})
	IsValid(value interface{}) (result bool, err *string)
}

type FieldMeta struct {
	Name  string
	Type  string
	Value interface{}
	Label string
}
