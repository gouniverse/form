package form

import (
	"strconv"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// == CLASS ===================================================================

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
	// BlockEditorOptions BlockEditorOptions
	Placeholder string
	Invisible   bool
	CustomInput hb.TagInterface
}

// == INTERFACES ==============================================================

var _ FieldInterface = (*Field)(nil)

// == IMPLEMENTATION OF FieldInterface ========================================

func (field *Field) GetID() string {
	return field.ID
}

func (field *Field) SetID(fieldID string) {
	field.ID = fieldID
}

func (field *Field) GetLabel() string {
	return field.Label
}

func (field *Field) SetLabel(fieldLabel string) {
	field.Label = fieldLabel
}

func (field *Field) GetHelp() string {
	return field.Help
}

func (field *Field) SetHelp(fieldHelp string) {
	field.Help = fieldHelp
}

func (field *Field) GetName() string {
	return field.Name
}

func (field *Field) SetName(fieldName string) {
	field.Name = fieldName
}

func (field *Field) GetOptions() []FieldOption {
	return field.Options
}

func (field *Field) SetOptions(fieldOptions []FieldOption) {
	field.Options = fieldOptions
}

func (field *Field) GetOptionsF() func() []FieldOption {
	return field.OptionsF
}

func (field *Field) SetOptionsF(fieldOptionsF func() []FieldOption) {
	field.OptionsF = fieldOptionsF
}

func (field *Field) GetRequired() bool {
	return field.Required
}

func (field *Field) SetRequired(fieldRequired bool) {
	field.Required = fieldRequired
}

func (field *Field) GetType() string {
	return field.Type
}

func (field *Field) SetType(fieldType string) {
	field.Type = fieldType
}

func (field *Field) GetValue() string {
	return field.Value
}

func (field *Field) SetValue(fieldValue string) {
	field.Value = fieldValue
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

// == METHODS ================================================================

func (field *Field) IsBlockEditor() bool {
	return field.Type == FORM_FIELD_TYPE_BLOCKEDITOR
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

		if field.Placeholder != "" {
			input.Placeholder(field.Placeholder)
		}

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
		input = field.fieldDateTime()
	}

	if field.IsImage() {
		input = field.fieldImage(fileManagerURL)
	}

	if field.IsHtmlArea() {
		input = field.fieldHtmlArea()
	}

	if field.IsBlockEditor() {
		input = field.fieldBlockEditor()
	}

	if field.IsSelect() {
		input = field.fieldSelect()
	}

	if field.IsTable() {
		input = field.fieldTable(fileManagerURL)
	}

	if field.IsTextArea() {
		input = field.fieldTextArea()
	}

	if field.IsReadonly() {
		// Selects are different. Readonly for selects does not work.
		// Disable and create a hidden field
		if field.IsSelect() {
			input.Attr("disabled", "disabled")
			input.Name(field.Name + "_Readonly")
		} else {
			input.Attr("readonly", "readonly")
		}
	}

	if field.IsDisabled() {
		input.Attr("disabled", "disabled")
	}

	return input
}

func (field *Field) fieldBlockEditor() *hb.Tag {
	textInputOnError := hb.NewTextArea().
		ID(field.ID).
		Class("form-control").
		Name(field.Name).
		Text(field.Value)

	if field.CustomInput == nil {
		return hb.Wrap().
			Child(hb.Div().
				Class("alert alert-danger").
				Text("CustomInput is nil")).
			Child(textInputOnError)
	}

	return hb.Wrap(field.CustomInput)
}

func (field *Field) fieldDateTime() *hb.Tag {
	input := hb.NewInput().
		ID(field.ID).
		Type(hb.TYPE_DATETIME).
		Class("form-control").
		Name(field.Name).
		Value(field.Value)
	// formGroupInput = hb.NewTag(`el-date-picker`).Attr("type", "datetime").Attr("v-model", "entityModel."+fieldName)
	// formGroupInput = hb.NewTag(`n-date-picker`).Attr("type", "datetime").Class("form-control").Attr("v-model", "entityModel."+fieldName)

	return input
}

func (field *Field) fieldImage(fileManagerURL string) *hb.Tag {
	image := hb.NewImage().
		Class(`img-fluid rounded-start`).
		Style(`margin-bottom: 15px;`).
		AttrIf(field.Value != "", `src`, field.Value).
		AttrIf(field.Value == "", `src`, `https://www.freeiconspng.com/uploads/no-image-icon-11.PNG`).
		Style(`width:100%;max-height:100px;`)

	textArea := hb.NewTextArea().
		ID(field.ID).
		Type(hb.TYPE_TEXT).
		Class("form-control").
		Style(`height:70px;`).
		Name(field.Name).
		Text(field.Value).
		Placeholder(field.Placeholder)

	fileManagerLink := lo.If(fileManagerURL != "", hb.NewHyperlink().Href(fileManagerURL).Target("_blank").Text("Browse")).
		Else(hb.Span().Text("The URL can be base64 encoded image URL"))

	input := hb.NewDiv().
		Class(`row g-3`).
		Style(`border: 1px solid silver;border-radius: 10px; margin-top: 0px; margin-left: 0px;margin-right: 0px;`).
		Child(hb.NewDiv().
			Class(`col-md-2`).
			Child(image),
		).
		Child(hb.NewDiv().
			Class(`col-md-10`).
			Child(textArea).
			Child(fileManagerLink),
		)

	return input
}

func (field *Field) fieldHtmlArea() *hb.Tag {
	textarea := hb.NewTextArea().
		ID(field.ID).
		Class("form-control").
		Name(field.Name).
		Text(field.Value)

	input := hb.NewWrap().
		Child(textarea).
		Child(hb.NewScript(field.TrumbowygScript()))
	// Child(hb.NewScript(`setTimeout(() => {initWysiwyg("` + field.ID + `")}, 100);`))

	return input
}

func (field *Field) fieldSelect() *hb.Tag {
	input := hb.NewSelect().
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
	return input
}

func (field *Field) fieldTable(fileManagerURL string) *hb.Tag {
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

	input := hb.NewWrap().Child(table)

	if field.TableOptions.RowAddButton != nil {
		input.AddChild(hb.NewDiv().Child(field.TableOptions.RowAddButton.Type(hb.TYPE_BUTTON)))
	}

	return input
}

func (field *Field) fieldTextArea() *hb.Tag {
	return hb.NewTextArea().
		ID(field.ID).
		Class("form-control").
		Name(field.Name).
		HTML(field.Value)
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

	formGroup := hb.NewDiv().
		Class("form-group mb-3")

	formGroupLabel := hb.NewLabel().
		HTML(fieldLabel).
		Class("form-label").
		ChildIf(
			field.Required,
			hb.NewSup().HTML("*").Class("text-danger ms-1"),
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

	if !field.IsHidden() {
		formGroupLabel.Attr("for", field.ID)
	}

	if field.Invisible {
		formGroup.Style("display:none;")
	}

	// Add help
	if field.Help != "" {
		formGroupHelp := hb.NewParagraph().Class("text-info").HTML(field.Help)
		formGroup.Child(formGroupHelp)
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

func (field *Field) clone() FieldInterface {
	fieldCopy := *field
	return &fieldCopy
}
