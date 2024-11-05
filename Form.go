package form

import (
	"github.com/gouniverse/hb"
)

type Form struct {
	id             string
	className      string
	fields         []FieldInterface
	fileManagerURL string
	method         string
	actionUrl      string

	// HTMX helpers
	hxPost   string
	hxTarget string
	hxSwap   string
}

func (form *Form) AddField(field FieldInterface) {
	form.fields = append(form.fields, field)
}

func (form *Form) GetFields() []FieldInterface {
	return form.fields
}

func (form *Form) SetFields(fields []FieldInterface) {
	form.fields = fields
}

func (form *Form) GetFileManagerURL() string {
	return form.fileManagerURL
}

func (form *Form) SetFileManagerURL(url string) {
	form.fileManagerURL = url
}

func (form *Form) Build() *hb.Tag {
	tags := []hb.TagInterface{}

	for _, field := range form.fields {
		tags = append(tags, field.BuildFormGroup(form.fileManagerURL))
	}

	hbForm := hb.Form()
	hbForm.Children(tags)
	hbForm.Method(form.method)

	if form.actionUrl != "" {
		hbForm.Action(form.actionUrl)
	}

	if form.id != "" {
		hbForm.ID(form.id)
	}

	if form.className != "" {
		hbForm.Class(form.className)
	}

	if form.hxPost != "" {
		hbForm.HxPost(form.hxPost)
	}

	if form.hxTarget != "" {
		hbForm.HxTarget(form.hxTarget)
	}

	if form.hxSwap != "" {
		hbForm.HxSwap(form.hxSwap)
	}

	return hbForm
}
