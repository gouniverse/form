package form

import (
	"strconv"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type Field struct {
	ID           string // automatic, if not assigned
	Type         string
	Name         string
	Label        string
	Help         string
	Options      []FieldOption
	OptionsF     func() []FieldOption
	Value        string
	Required     bool
	Readonly     bool
	Disabled     bool
	TableOptions TableOptions
}

type TableColumn struct {
	Label string
	Width int
}
type TableOptions struct {
	Header          []TableColumn
	Rows            [][]Field
	RowAddButton    *hb.Tag
	RowDeleteButton *hb.Tag
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

func (field *Field) IsTable() bool {
	return field.Type == FORM_FIELD_TYPE_TABLE
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

func (field *Field) fieldInput(fileManagerURL string) *hb.Tag {
	if field.IsRaw() {
		return hb.NewHTML(field.Value)
	}

	if field.ID == "" {
		field.ID = "id_" + uid.HumanUid()
	}

	input := hb.NewTag(``) // no tag by default

	if field.IsDate() ||
		field.IsHidden() ||
		field.IsPassword() ||
		field.IsString() ||
		field.IsNumber() {

		input = hb.NewInput().
			ID(field.ID).
			Class("form-control").
			Name(field.Name).
			Value(field.Value)

		if field.IsDate() {
			input.Type(hb.TYPE_DATE)
		}

		if field.IsHidden() {
			input.Type(hb.TYPE_HIDDEN)
		}

		if field.IsNumber() {
			input.Type(hb.TYPE_NUMBER)
		}

		if field.IsPassword() {
			input.Type(hb.TYPE_PASSWORD)
		}

		if field.IsString() {
			input.Type(hb.TYPE_TEXT)
		}
	}

	if field.IsDateTime() {
		input = hb.NewInput().
			ID(field.ID).
			Type(hb.TYPE_DATETIME).
			Class("form-control").
			Name(field.Name).
			Value(field.Value)
		// formGroupInput = hb.NewTag(`el-date-picker`).Attr("type", "datetime").Attr("v-model", "entityModel."+fieldName)
		// formGroupInput = hb.NewTag(`n-date-picker`).Attr("type", "datetime").Class("form-control").Attr("v-model", "entityModel."+fieldName)
	}

	if field.IsImage() {
		input = hb.NewDiv().Children([]hb.TagInterface{
			hb.NewImage().
				AttrIf(field.Value != "", `src`, field.Value).
				AttrIf(field.Value == "", `src`, `https://www.freeiconspng.com/uploads/no-image-icon-11.PNG`).
				Style(`width:200px;`),

			bs.InputGroup().Children([]hb.TagInterface{
				hb.NewInput().
					ID(field.ID).
					Type(hb.TYPE_URL).
					Class("form-control").
					Name(field.Name).
					Value(field.Value),
				hb.If(fileManagerURL != "", bs.InputGroupText().Children([]hb.TagInterface{
					hb.NewHyperlink().HTML("Browse").Href(fileManagerURL).Target("_blank"),
				})),
			}),
		})
	}

	if field.IsHtmlArea() {
		input = hb.NewWrap().Child(
			hb.NewTextArea().
				ID(field.ID).
				Class("form-control").
				Name(field.Name).
				Text(field.Value)).
			Child(hb.NewScript(field.TrumbowygScript()))
		// Child(hb.NewScript(`setTimeout(() => {initWysiwyg("` + field.ID + `")}, 100);`))
	}

	if field.IsSelect() {
		input = hb.NewSelect().
			ID(field.ID).
			Name(field.Name).
			Class("form-select")

		for _, opt := range field.Options {
			option := hb.NewOption().Value(opt.Key).HTML(opt.Value)
			option.AttrIf(field.Value == opt.Key, "selected", "selected")
			input.AddChild(option)
		}
		if field.OptionsF != nil {
			for _, opt := range field.OptionsF() {
				option := hb.NewOption().Value(opt.Key).HTML(opt.Value)
				option.AttrIf(field.Value == opt.Key, "selected", "selected")
				input.AddChild(option)
			}
		}
	}

	if field.IsTable() {
		header := hb.NewThead()
		if field.TableOptions.RowDeleteButton != nil {
			th := hb.NewTH().HTML("#").Style("width:1px;")
			header.AddChild(th)
		}
		for _, v := range field.TableOptions.Header {
			th := hb.NewTH().HTML(v.Label)
			if v.Width != 0 {
				th.Style("width:" + strconv.Itoa(v.Width) + "px")
			}
			header.AddChild(th)
		}

		rows := hb.NewTbody()
		for rowIndex, rowFields := range field.TableOptions.Rows {
			tr := hb.NewTR().Data("row-index", utils.ToString(rowIndex))
			if field.TableOptions.RowDeleteButton != nil {
				deleteButton := field.TableOptions.RowDeleteButton.
					Type(hb.TYPE_BUTTON).
					Data("row-index", utils.ToString(rowIndex))
				td := hb.NewTH().Child(deleteButton)
				tr.AddChild(td)
			}
			for _, rowField := range rowFields {
				td := hb.NewTD().Child(rowField.fieldInput(fileManagerURL))
				tr.AddChild(td)
			}
			rows.AddChild(tr)
		}
		table := hb.NewTable().
			ID(field.ID).
			Class("table table-striped table-hover mb-0").
			Child(header).
			Child(rows)

		input = hb.NewWrap().Child(table)

		if field.TableOptions.RowAddButton != nil {
			input.AddChild(hb.NewDiv().Child(field.TableOptions.RowAddButton.Type(hb.TYPE_BUTTON)))
		}
	}

	if field.IsTextArea() {
		input = hb.NewTextArea().
			ID(field.ID).
			Name(field.Name).
			Class("form-control").
			HTML(field.Value)
	}

	if field.IsReadonly() {
		// Selects are different. Readonly for selects does not work.
		// Disable and create a hidden field
		if field.IsSelect() {
			input.Attr("disabled", "disabled")
			input.Name(field.Name + "_Readonly")
			// hiddenInput = hb.NewInput().
			// 	Class("form-control").
			// 	Name(field.Name).
			// 	Value(field.Value).
			// 	Type(hb.TYPE_HIDDEN)
		} else {
			input.Attr("readonly", "readonly")
		}
	}

	if field.IsDisabled() {
		input.Attr("disabled", "disabled")
	}

	return input
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

	formGroup := hb.NewDiv().Class("form-group mb-3")

	formGroupLabel := hb.NewLabel().
		HTML(fieldLabel).
		Class("form-label").
		ChildIf(
			field.Required,
			hb.NewSup().HTML("*").Class("text-danger ml-1"),
		)

	// Hidden input
	hiddenInput := hb.NewTag(``) // special case, no tag by default

	if field.IsReadonly() && field.IsSelect() {
		hiddenInput = hb.NewInput().
			Class("form-control").
			Name(field.Name).
			Value(field.Value).
			Type(hb.TYPE_HIDDEN)
	}

	if !field.IsHidden() {
		formGroup.Child(formGroupLabel)
	}
	formGroup.Child(field.fieldInput(fileManagerURL))
	formGroup.Child(hiddenInput)

	// Add help
	if field.Help != "" {
		formGroupHelp := hb.NewParagraph().Class("text-info").HTML(field.Help)
		formGroup.AddChild(formGroupHelp)
	}

	return formGroup
}

func (field *Field) TrumbowygScript() string {
	fieldConfig, found := lo.Find(field.Options, func(fieldOption FieldOption) bool {
		return fieldOption.Key == "config"
	})

	config := "null"

	if found {
		config = fieldConfig.Value
	}

	return `
if (!window.trumbowigConfig) {
	window.trumbowigConfig = {
		btns: [
			['formatting'],
			['strong', 'em', 'del'],
			['superscript', 'subscript'],
			['link','justifyLeft','justifyRight','justifyCenter','justifyFull'],
			['unorderedList', 'orderedList'],
			['removeformat'],
			['undo', 'redo'],
			['horizontalRule'],
			['fullscreen'],
		],
		autogrow: true,
		removeformatPasted: true,
		tagsToRemove: ['script', 'link', 'embed', 'iframe', 'input'],
		tagsToKeep: ['hr', 'img', 'i'],
		autogrowOnEnter: true,
		linkTargets: ['_blank'],
	};

	function initWysiwyg(textareaID, config) {
	    var editorConfig = config || window.trumbowigConfig;
		$('#' + textareaID).trumbowyg(editorConfig);
	}
}

setTimeout(() => {
	initWysiwyg("` + field.ID + `", ` + config + `);	
}, 200);
`
}
