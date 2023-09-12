package form

import (
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
)

type Field struct {
	Type     string
	Name     string
	Label    string
	Help     string
	Options  []FieldOption
	OptionsF func() []FieldOption
	Value    string
	Required bool
	Readonly bool
	Disabled bool
}

func (field *Field) IsDate() bool {
	return field.Type == FORM_FIELD_TYPE_DATE
}

func (field *Field) IsDateTime() bool {
	return field.Type == FORM_FIELD_TYPE_DATETIME
}

func (field *Field) IsImage() bool {
	return field.Type == FORM_FIELD_TYPE_IMAGE
}

func (field *Field) IsHidden() bool {
	return field.Type == FORM_FIELD_TYPE_HIDDEN
}

func (field *Field) IsHtmlArea() bool {
	return field.Type == FORM_FIELD_TYPE_HTMLAREA
}

func (field *Field) IsNumber() bool {
	return field.Type == FORM_FIELD_TYPE_NUMBER
}

func (field *Field) IsPassword() bool {
	return field.Type == FORM_FIELD_TYPE_PASSWORD
}

func (field *Field) IsSelect() bool {
	return field.Type == FORM_FIELD_TYPE_SELECT
}

func (field *Field) IsString() bool {
	return field.Type == FORM_FIELD_TYPE_STRING
}

func (field *Field) IsTextArea() bool {
	return field.Type == FORM_FIELD_TYPE_TEXTAREA
}

func (field *Field) IsReadonly() bool {
	return field.Readonly
}

func (field *Field) IsDisabled() bool {
	return field.Disabled
}

func (field *Field) IsRequired() bool {
	return field.Required
}

func (field *Field) IsRaw() bool {
	return field.Type == FORM_FIELD_TYPE_RAW
}

func (field *Field) BuildFormGroup(fileManagerURL string) *hb.Tag {
	if field.IsRaw() {
		return hb.NewHTML(field.Value)
	}

	fieldName := field.Name
	fieldLabel := field.Label
	if fieldLabel == "" {
		fieldLabel = fieldName
	}

	formGroup := hb.NewDiv().Class("form-group mt-3")

	formGroupLabel := hb.NewLabel().
		HTML(fieldLabel).
		Class("form-label").
		ChildIf(
			field.Required,
			hb.NewSup().HTML("*").Class("text-danger ml-1"),
		)

	formGroupInput := hb.NewTag(``) // no tag by default
	hiddenInput := hb.NewTag(``)    // special case, no tag by default

	if field.IsDate() ||
		field.IsHidden() ||
		field.IsPassword() ||
		field.IsString() ||
		field.IsNumber() {
		formGroupInput = hb.NewInput().
			Class("form-control").
			Name(fieldName).
			Value(field.Value)

		if field.IsDate() {
			formGroupInput.Type(hb.TYPE_DATE)
		}

		if field.IsHidden() {
			formGroupInput.Type(hb.TYPE_HIDDEN)
		}

		if field.IsNumber() {
			formGroupInput.Type(hb.TYPE_NUMBER)
		}

		if field.IsPassword() {
			formGroupInput.Type(hb.TYPE_PASSWORD)
		}

		if field.IsString() {
			formGroupInput.Type(hb.TYPE_TEXT)
		}
	}

	if field.IsDateTime() {
		formGroupInput = hb.NewInput().
			Type(hb.TYPE_DATETIME).
			Class("form-control").
			Name(fieldName).
			Value(field.Value)
		// formGroupInput = hb.NewTag(`el-date-picker`).Attr("type", "datetime").Attr("v-model", "entityModel."+fieldName)
		// formGroupInput = hb.NewTag(`n-date-picker`).Attr("type", "datetime").Class("form-control").Attr("v-model", "entityModel."+fieldName)
	}

	if field.IsImage() {
		formGroupInput = hb.NewDiv().Children([]*hb.Tag{
			hb.NewImage().
				AttrIf(field.Value != "", `src`, field.Value).
				AttrIf(field.Value == "", `src`, `https://www.freeiconspng.com/uploads/no-image-icon-11.PNG`).
				Style(`width:200px;`),

			bs.InputGroup().Children([]*hb.Tag{
				hb.NewInput().
					Type(hb.TYPE_URL).
					Class("form-control").
					Name(fieldName).
					Value(field.Value),
				hb.If(fileManagerURL != "", bs.InputGroupText().Children([]*hb.Tag{
					hb.NewHyperlink().HTML("Browse").Href(fileManagerURL).Target("_blank"),
				})),
			}),
		})
	}

	if field.IsHtmlArea() {
		formGroupInput = hb.NewTag("trumbowyg").
			Name(fieldName).
			Attr("v-model", "entityModel."+fieldName).
			Attr(":config", "trumbowigConfig").
			Class("form-control")
	}

	if field.IsSelect() {
		formGroupInput = hb.NewSelect().
			Name(fieldName).
			Class("form-select")

		for _, opt := range field.Options {
			option := hb.NewOption().Value(opt.Key).HTML(opt.Value)
			option.AttrIf(field.Value == opt.Key, "selected", "selected")
			formGroupInput.AddChild(option)
		}
		if field.OptionsF != nil {
			for _, opt := range field.OptionsF() {
				option := hb.NewOption().Value(opt.Key).HTML(opt.Value)
				option.AttrIf(field.Value == opt.Key, "selected", "selected")
				formGroupInput.AddChild(option)
			}
		}
	}

	if field.IsTextArea() {
		formGroupInput = hb.NewTextArea().
			Name(fieldName).
			Class("form-control").
			HTML(field.Value)
	}

	if field.IsReadonly() {
		// Selects are different. Readonly for selects does not work.
		// Disable and create a hidden field
		if field.IsSelect() {
			formGroupInput.Attr("disabled", "disabled")
			formGroupInput.Name(field.Name + "_Readonly")
			hiddenInput = hb.NewInput().
				Class("form-control").
				Name(field.Name).
				Value(field.Value).
				Type(hb.TYPE_HIDDEN)
		} else {
			formGroupInput.Attr("readonly", "readonly")
		}
	}

	if field.IsDisabled() {
		formGroupInput.Attr("disabled", "disabled")
	}

	if !field.IsHidden() {
		formGroup.Child(formGroupLabel)
	}
	formGroup.Child(formGroupInput)
	formGroup.Child(hiddenInput)

	// Add help
	if field.Help != "" {
		formGroupHelp := hb.NewParagraph().Class("text-info").HTML(field.Help)
		formGroup.AddChild(formGroupHelp)
	}

	return formGroup
}
