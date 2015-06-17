package forms

import (
	"encoding/json"
)

type Form struct {
	Name   string
	Label  string
	Fields map[string]Field
}

type FormMeta struct {
	Name   string
	Label  string
	Fields map[string]FieldMeta
}

func (form *Form) Unmarshal(jsonBytes []byte) {

	meta := FormMeta{}

	err := json.Unmarshal(jsonBytes, &meta)

	if err != nil {
		panic(err)
	}

	form.Name = meta.Name
	form.Label = meta.Label
	form.Fields = make(map[string]Field)

	if len(meta.Fields) > 0 {
		for _, item := range meta.Fields {

			field, err := FieldsFactory.CreateField(item)

			if err != nil {
				panic(err)
			}

			form.Fields[(*field).GetName()] = *field
		}
	}
}

func (form Form) Marshal() []byte {

	result, err := json.Marshal(form)

	if err != nil {
		panic(err)
	}

	return result
}
