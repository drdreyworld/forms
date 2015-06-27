package forms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Form struct {
	Name    string
	Label   string
	Fields  map[string]Field
	Buttons map[string]Field
}

type FormMeta struct {
	Name    string
	Label   string
	Fields  map[string]FieldMeta
	Buttons map[string]FieldMeta
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (form *Form) IsValid(r *http.Request) bool {
	result := true
	for name, field := range form.Fields {
		isValid, _ := field.IsValid(r.PostFormValue(name))
		result = isValid && result
	}
	return result
}

func (form *Form) GetValues() map[string]interface{} {
	result := make(map[string]interface{})
	for name, field := range form.Fields {
		result[name] = field.GetValue()
	}
	return result
}

func (form *Form) RenderField(field Field) template.HTML {

	name := field.GetType()
	path := "/Users/andrey/GoProjects/src/github.com/drdreyworld/news/templates/view/admin/form-edit"

	filename, err := filepath.Abs(fmt.Sprintf("%s/%s.html", path, name))
	checkErr(err)

	tmpl, err := template.New("").ParseFiles(filename)
	checkErr(err)

	buffer := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buffer, name, field)
	checkErr(err)

	return template.HTML(buffer.String())
}

func (form *Form) CreateFromMeta(meta FormMeta) {
	form.Name = meta.Name
	form.Label = meta.Label
	form.Fields = make(map[string]Field)
	form.Buttons = make(map[string]Field)

	if len(meta.Fields) > 0 {
		for _, item := range meta.Fields {

			field, err := FieldsFactory.CreateField(item)

			if err != nil {
				panic(err)
			}

			form.Fields[(*field).GetName()] = *field
		}
	}
	if len(meta.Buttons) > 0 {
		for _, item := range meta.Buttons {

			field, err := FieldsFactory.CreateField(item)

			if err != nil {
				panic(err)
			}

			form.Buttons[(*field).GetName()] = *field
		}
	}
}

func (form *Form) Unmarshal(jsonBytes []byte) {

	meta := FormMeta{}

	checkErr(json.Unmarshal(jsonBytes, &meta))

	form.CreateFromMeta(meta)
}

func (form Form) Marshal() []byte {
	result, err := json.Marshal(form)
	checkErr(err)

	return result
}
