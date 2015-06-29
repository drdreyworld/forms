# Пакет forms

Пакет для работы с web-формами.

Предполагаемый алгоритм использования:

1. конфиг формы берем из json или из подготовленной структуры FormMeta
2. в конфиге полностью описана структуры формы, включая поля, кнопки, валидаторы
3. создаем структуру Form
4. устанавливаем значения элементов
5. рисуем форму
6. получаем POST request
7. валидируем полученные значения
8. сохраняем нормализованные значения или денормализованную структуру формы (FormMeta)


# Шаблоны

Предполагается, что для форм будут подготовлены два пакета шаблонов:

1. Шаблоны для режима "редактирование"
2. Шаблоны для режима "представления" заполненной формы


Полный пример использования: https://github.com/drdreyworld/news

*Пример создания формы из подготовленной структуры FormMeta*

```
func (item NewsItem) GetForm() forms.FormMeta {
	result := forms.FormMeta{
		Name:  "AdminNewsForm",
		Label: "Редактировать новость",
		Fields: map[string]fields.FieldMeta{
			"title": fields.FieldMeta{
				Name:  "title",
				Type:  "text",
				Label: "Заголовок",
				Value: item.title,
			},
			"is_published": fields.FieldMeta{
				Name:  "is_published",
				Type:  "checkbox",
				Label: "Новость опубликована",
				Value: item.is_published,
			},
			"image": fields.FieldMeta{
				Name:  "image",
				Type:  "text",
				Label: "URL картинки",
				Value: item.image,
			},
			"anounce": fields.FieldMeta{
				Name:  "anounce",
				Type:  "textarea",
				Label: "Анонс",
				Value: item.anounce,
			},
			"content": fields.FieldMeta{
				Name:  "content",
				Type:  "textarea",
				Label: "Полный текст новости",
				Value: item.anounce,
			},
		},
		Buttons: map[string]fields.FieldMeta{
			"save": fields.FieldMeta{
				Name:  "save",
				Type:  "submit",
				Label: "Сохранить",
				Value: "save",
			},
		},
	}
	return result
}
```

*Пример создания формы*

```
func (ctrl *AdminController) CreateItemGet(r *http.Request, params httprouter.Params) interface{} {
	item := model.NewsItem{}
	form := new(forms.Form)
	form.CreateFromMeta(item.GetForm())
	form.SetTemplatesPath(ctrl.config.Paths.Templates + "/view/admin/form-edit/")

	return map[string]interface{}{
		"item": item,
		"form": form,
	}
}
```

*Пример шаблона формы*

```
{{define "view/admin/form-edit"}}
{{ $form := .form }}
<form method="post" class="form-horizontal">
	<div class="panel panel-default">
		<div class="panel-heading">{{ $form.Label }}</div>
		<div class="panel-body">
			{{ range $key, $field := $form.Fields }}
				<div class="panel-group">
					<div class="form-group">
						<label class="col-sm-2 control-label">{{ $field.GetLabel }}</label>
						<div class="col-sm-10">
							{{ $form.RenderField $field }}
						</div>
					</div>
				</div>
			{{ end }}
		</div>
		<div class="panel-footer" style="text-align: right">
			{{ range $key, $button := $form.Buttons }}
				{{ $form.RenderField $button }}
			{{ end }}
		</div>
	</div>
</form>

{{end}}
```


