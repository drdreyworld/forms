/*
	Package forms implements work with web-forms.

	Usage

	Create form from struct

		meta := forms.FormMeta{
			Name:  "AdminNewsForm",
			Label: "Редактировать новость",
			Fields: fields.FieldsMeta{
				fields.FieldMeta{
					Name:  "title",
					Type:  "text",
					Label: "Заголовок",
					Value: item.title,
					Order: 0,
				},
				fields.FieldMeta{
					Name:  "is_published",
					Type:  "checkbox",
					Label: "Новость опубликована",
					Value: item.is_published,
					Order: 1,
				},
				fields.FieldMeta{
					Name:  "image",
					Type:  "text",
					Label: "URL картинки",
					Value: item.image,
					Order: 2,
				},
				fields.FieldMeta{
					Name:  "anounce",
					Type:  "textarea",
					Label: "Анонс",
					Value: item.anounce,
					Order: 3,
				},
				fields.FieldMeta{
					Name:  "content",
					Type:  "textarea",
					Label: "Полный текст новости",
					Value: item.anounce,
					Order: 4,
				},
			},
			Buttons: fields.FieldsMeta{
				fields.FieldMeta{
					Name:  "save",
					Type:  "submit",
					Label: "Сохранить",
					Value: "save",
				},
			},
		};

		form := new(forms.Form)
		form.CreateFromMeta(meta)
		form.SetTemplatesPath("/www/view/admin/form-edit/")

*/
package forms
