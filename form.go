package forms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/drdreyworld/forms/fields"
	"github.com/drdreyworld/forms/validators"
	"github.com/drdreyworld/webapp"
	"html/template"
	"net/http"
	"path/filepath"
	"sort"
)

type Form struct {
	Name          string
	Label         string
	Fields        fields.Fields
	Buttons       fields.Fields
	TemplatesPath string
}

type FormMeta struct {
	Name    string
	Label   string
	Fields  fields.FieldsMeta
	Buttons fields.FieldsMeta
}

func (form *Form) SetTemplatesPath(path string) {
	form.TemplatesPath = path
}

func (form *Form) GetTemplatesPath() string {
	return form.TemplatesPath
}

func (form *Form) IsValid(r *http.Request) bool {
	result := true
	for _, field := range form.Fields {
		result = field.IsValid(r.PostFormValue(field.GetName())) && result
	}
	return result
}

func (form *Form) GetValues() map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range form.Fields {
		result[field.GetName()] = field.GetValue()
	}
	return result
}

func (form *Form) GetValue(name string) interface{} {
	if field, ok := form.Fields.GetField(name); ok {
		return field.GetValue()
	}
	return nil
}

func (form *Form) RenderField(field fields.Field) template.HTML {

	name := field.GetType()
	path := form.GetTemplatesPath()

	filename, err := filepath.Abs(fmt.Sprintf("%s/%s.html", path, name))
	webapp.Panic(err)

	tmpl, err := template.New("").ParseFiles(filename)
	webapp.Panic(err)

	buffer := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buffer, name, field)
	webapp.Panic(err)

	return template.HTML(buffer.String())
}

func (form *Form) CreateFromMeta(meta FormMeta) {
	form.Name = meta.Name
	form.Label = meta.Label
	form.Fields = make(fields.Fields, 0, len(meta.Fields))
	form.Buttons = make(fields.Fields, 0, len(meta.Buttons))

	if len(meta.Fields) > 0 {
		sort.Sort(meta.Fields)
		for _, item := range meta.Fields {
			field, err := fields.Factory.CreateField(item)
			webapp.Panic(err)

			fieldValidators := make(validators.Validators, 0, len(item.Validators))

			for _, validatorMeta := range item.Validators {
				if fieldValidator, err := validators.Factory.CreateValidator(validatorMeta); err == nil {
					fieldValidators = append(fieldValidators, fieldValidator)
				}
			}
			field.SetValidators(fieldValidators)

			form.Fields = append(form.Fields, field)
		}
	}
	if len(meta.Buttons) > 0 {
		sort.Sort(meta.Buttons)
		for _, item := range meta.Buttons {
			field, err := fields.Factory.CreateField(item)
			webapp.Panic(err)
			form.Buttons = append(form.Buttons, field)
		}
	}
}

func UnmarshalFormMeta(jsonBytes []byte) FormMeta {
	meta := FormMeta{}
	webapp.Panic(json.Unmarshal(jsonBytes, &meta))
	return meta
}

func (form *Form) Unmarshal(jsonBytes []byte) {
	form.CreateFromMeta(UnmarshalFormMeta(jsonBytes))
}

func (form Form) Marshal() []byte {
	result, err := json.Marshal(form)
	webapp.Panic(err)
	return result
}
