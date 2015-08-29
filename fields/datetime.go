package fields

import "time"

func init() {
	Factory.Register("datetime", new(DateTime))
}

type DateTime struct {
	BaseField
	Value string
}

func (prototype DateTime) Create(meta FieldMeta) Field {
	field := new(DateTime)
	field.Type = "datetime"
	field.SetName(meta.Name)
	field.SetLabel(meta.Label)
	field.SetValue(meta.Value)
	return field
}

func (field DateTime) GetValueInHumanFormat() interface{} {
	if len(field.Value) == 0 || field.Value == "0000-00-00 00:00:00" {
		return ""
	}

	if t, err := time.Parse("2006-01-02 15:04:05", field.Value); err == nil {
		return t.Format("02.01.2006 15:04")
	}
	return field.Value
}

func (field DateTime) GetValue() interface{} {
	return field.Value
}

func (field *DateTime) SetValue(value interface{}) (ok bool) {
	field.Value, ok = value.(string)
	return ok
}

func (field *DateTime) IsValid(value interface{}) (result bool) {
	if result = field.SetValue(value); result {
		for _, validator := range field.Validators {
			if result = validator.IsValid(value); !result {
				field.Error = validator.GetError()
				break
			}
		}
		if result && len(field.Value) > 0 {
			t, _ := time.Parse("02.01.2006 15:04", field.Value)
			field.Value = t.Format("2006-01-02 15:04:05")
		}
	}
	return result
}
