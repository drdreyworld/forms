package fields

type Field interface {
	Create(meta FieldMeta) Field
	GetType() string
	GetLabel() string
	SetLabel(label string)
	GetName() string
	SetName(name string)
	GetValue() interface{}
	SetValue(value interface{})
	IsValid(value interface{}) (result bool, err *string)
	GetOrder() int
	SetOrder(order int)
}

type Fields []Field

func (fields Fields) Len() int {
	return len(fields)
}

func (fields Fields) Less(i, j int) bool {
	return fields[i].GetOrder() < fields[j].GetOrder()
}

func (fields Fields) Swap(i, j int) {
	fields[i], fields[j] = fields[j], fields[i]
}

type FieldMeta struct {
	Name  string
	Type  string
	Value interface{}
	Label string
	Order int
}

type FieldsMeta []FieldMeta

func (fields FieldsMeta) Len() int {
	return len(fields)
}

func (fields FieldsMeta) Less(i, j int) bool {
	return fields[i].Order < fields[j].Order
}

func (fields FieldsMeta) Swap(i, j int) {
	fields[i], fields[j] = fields[j], fields[i]
}
