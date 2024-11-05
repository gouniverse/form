package form

import "net/http"

type FormOptions struct {
	ActionURL      string           // optional
	ClassName      string           // optional
	ID             string           // optional
	Fields         []FieldInterface // optional
	FileManagerURL string           // optional
	Method         string           // optional

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
