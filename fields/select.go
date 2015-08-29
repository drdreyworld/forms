package fields

func init() {
	Factory.Register("select", new(Select))
}

type Select struct {
	BaseField
	Value        ValueOption
	ValueOptions ValueOptions
}

func (prototype Select) Create(meta FieldMeta) Field {
	field := new(Select)
	field.Type = "select"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field Select) GetValue() interface{} {
	return field.Value
}

func (field *Select) SetValue(value interface{}) (ok bool) {

	switch value.(type) {
	case string:
		field.Value, ok = field.ValueOptions.GetOptionByValue(value)
	case ValueOption:
		field.Value, ok = value.(ValueOption)
	case map[string]interface{}:
		if _, ok := value.(map[string]interface{})["Value"]; ok {
			return field.SetValue(ValueOption{
				Value: value.(map[string]interface{})["Value"],
				Title: value.(map[string]interface{})["Title"],
			})
		}
	}
	return ok
}

func (field *Select) IsValid(value interface{}) (result bool) {
	if len(field.ValueOptions) < 1 {
		field.Error = "Список для выбора пуст"
		return false
	}

	if result = field.SetValue(value); result {
		for _, validator := range field.Validators {
			if result = validator.IsValid(value); !result {
				field.Error = validator.GetError()
				break
			}
		}
	} else {
		field.Error = "Не допустимый вариант"
	}

	return result
}

func (field *Select) SetValueOptions(options ValueOptions) {
	field.ValueOptions = options
	field.SetValue(field.Value)
}

func (field Select) GetValueOptions() ValueOptions {
	return field.ValueOptions
}
