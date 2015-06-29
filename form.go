package forms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/drdreyworld/forms/fields"
	"html/template"
	"net/http"
	"path/filepath"
)

type Form struct {
	Name          string
	Label         string
	Fields        map[string]fields.Field
	Buttons       map[string]fields.Field
	TemplatesPath string
}

type FormMeta struct {
	Name    string
	Label   string
	Fields  map[string]fields.FieldMeta
	Buttons map[string]fields.FieldMeta
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (form *Form) SetTemplatesPath(path string) {
	form.TemplatesPath = path
}

func (form *Form) GetTemplatesPath() string {
	return form.TemplatesPath
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

func (form *Form) RenderField(field fields.Field) template.HTML {

	name := field.GetType()
	path := form.GetTemplatesPath()

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
	form.Fields = make(map[string]fields.Field)
	form.Buttons = make(map[string]fields.Field)

	if len(meta.Fields) > 0 {
		for _, item := range meta.Fields {

			field, err := fields.Factory.CreateField(item)

			if err != nil {
				panic(err)
			}

			form.Fields[(*field).GetName()] = *field
		}
	}
	if len(meta.Buttons) > 0 {
		for _, item := range meta.Buttons {

			field, err := fields.Factory.CreateField(item)

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
