package form

import (
	"net/http"

	"github.com/gouniverse/hb"
)

type Form struct {
	id             string
	className      string
	fields         []Field
	fileManagerURL string
	method         string
	actionUrl      string

	// HTMX helpers
	hxPost   string
	hxTarget string
	hxSwap   string
}

type FormOptions struct {
	ActionURL      string  // optional
	ClassName      string  // optional
	ID             string  // optional
	Fields         []Field // optional
	FileManagerURL string  // optional
	Method         string  // optional

	// HTMX helpers
	HxPost   string // optional
	HxTarget string // optional
	HxSwap   string // optional

}

func NewForm(opts FormOptions) *Form {
	form := &Form{}
	form.fields = opts.Fields
	form.fileManagerURL = opts.FileManagerURL
	form.method = opts.Method
	if form.method == "" {
		form.method = http.MethodPost
	}
	form.actionUrl = opts.ActionURL
	form.id = opts.ID
	form.className = opts.ClassName
	form.hxPost = opts.HxPost
	form.hxTarget = opts.HxTarget
	form.hxSwap = opts.HxSwap
	return form
}

func (form *Form) AddField(field Field) {
	form.fields = append(form.fields, field)
}

func (form *Form) GetFields() []Field {
	return form.fields
}

func (form *Form) SetFields(fields []Field) {
	form.fields = fields
}

func (form *Form) GetFileManagerURL() string {
	return form.fileManagerURL
}

func (form *Form) SetFileManagerURL(url string) {
	form.fileManagerURL = url
}

func (form *Form) Build() *hb.Tag {
	tags := []*hb.Tag{}

	for _, field := range form.fields {
		tags = append(tags, field.BuildFormGroup(form.fileManagerURL))
	}

	hbForm := hb.NewForm()
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
