package form

import (
	"strings"

	"github.com/gouniverse/hb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// Notes: still under development
// How it should work:
// - should accept a group of fields
// - should accept a list of values
// - should repeat the groups for each value in the list
// - should allow the user to edit the fields
// - should allow the user to add a new group via HTMX url
// - should allow the user to remove a group via HTMX url

// == CLASS ===================================================================

type fieldRepeater struct {
	form                *Form
	repeaterAddUrl      string
	repeaterMoveUpUrl   string
	repeaterMoveDownUrl string
	repeaterRemoveUrl   string
	fieldHelp           string
	fieldID             string
	fieldLabel          string
	fieldType           string
	fieldName           string
	fieldValue          string
	fields              []FieldInterface
	values              []map[string]string
}

// == INTERFACE ===============================================================

var _ FieldInterface = (*fieldRepeater)(nil)

// == IMPLEMENTATION OF FieldInterface ========================================

func (field *fieldRepeater) clone() FieldInterface {
	fieldCopy := *field
	return &fieldCopy
}

func (field *fieldRepeater) GetID() string {
	return field.fieldID
}

func (field *fieldRepeater) SetID(fieldID string) {
	field.fieldID = fieldID
}

func (field *fieldRepeater) GetLabel() string {
	return field.fieldLabel
}

func (field *fieldRepeater) SetLabel(fieldLabel string) {
	field.fieldLabel = fieldLabel
}

func (field *fieldRepeater) GetHelp() string {
	return field.fieldHelp
}

func (field *fieldRepeater) SetHelp(fieldHelp string) {
	field.fieldHelp = fieldHelp
}

func (field *fieldRepeater) GetName() string {
	return field.fieldName
}

func (field *fieldRepeater) SetName(fieldName string) {
	field.fieldName = fieldName
}

// GetOptions is not really needed for the repeater field
func (field *fieldRepeater) GetOptions() []FieldOption {
	return []FieldOption{}
}

// SetOptions is not really needed for the repeater field
func (field *fieldRepeater) SetOptions(fieldOptions []FieldOption) {

}

// GetOptionsF is not really needed for the repeater field
func (field *fieldRepeater) GetOptionsF() func() []FieldOption {
	return func() []FieldOption {
		return []FieldOption{}
	}
}

// SetOptionsF is not really needed for the repeater field
func (field *fieldRepeater) SetOptionsF(fieldOptionsF func() []FieldOption) {

}

// GetRequired is always false for the repeater field
func (field *fieldRepeater) GetRequired() bool {
	return false
}

// SetRequired is not really needed for the repeater field
func (field *fieldRepeater) SetRequired(fieldRequired bool) {
}

func (field *fieldRepeater) GetType() string {
	return field.fieldType
}

func (field *fieldRepeater) SetType(fieldType string) {
	field.fieldType = fieldType
}

func (field *fieldRepeater) GetValue() string {
	return field.fieldValue
}

func (field *fieldRepeater) SetValue(fieldValue string) {
	field.fieldValue = fieldValue
}

func (field *fieldRepeater) BuildFormGroup(fileManagerURL string) *hb.Tag {
	if field.repeaterAddUrl == "" {
		return hb.Div().Class("alert alert-danger").Text("Form Error. Repeater " + field.GetName() + " has no repeaterAddUrl")
	}

	if field.repeaterRemoveUrl == "" {
		return hb.Div().Class("alert alert-danger").Text("Form Error. Repeater " + field.GetName() + " has no repeaterRemoveUrl")
	}

	if !strings.Contains(field.repeaterRemoveUrl, "?") {
		field.repeaterRemoveUrl += `?`
	}

	fieldName := field.GetName()

	fieldLabel := field.GetLabel()
	if fieldLabel == "" {
		fieldLabel = fieldName
	}

	buttonAdd := hb.NewButton().
		Child(hb.I().Class("bi bi-plus")).
		HTML(" Add new").
		Class("btn btn-sm btn-primary ms-3 float-end").
		HxInclude("#" + field.form.id).
		HxPost(field.repeaterAddUrl).
		HxTarget("#" + field.form.id)

	formGroupLabel := hb.NewLabel().
		HTML(fieldLabel).
		Class("form-label").
		ChildIf(
			field.GetRequired(),
			hb.NewSup().HTML("*").Class("text-danger ms-1"),
		).
		Child(buttonAdd)

	cards := hb.Wrap()

	for index, mapKeyValue := range field.values {
		children := lo.Map(field.fields, func(field FieldInterface, index int) hb.TagInterface {
			clonedField := field.clone()

			clonedField.SetID(clonedField.GetID() + `_` + cast.ToString(index))
			fieldName := clonedField.GetName()
			fieldRepeaterValue := lo.ValueOr(mapKeyValue, fieldName, "")
			fieldRepeaterName := fieldName + `[]` // + `[` + cast.ToString(index) + `]`

			clonedField.SetName(fieldRepeaterName)
			clonedField.SetValue(fieldRepeaterValue)

			return clonedField.BuildFormGroup(fileManagerURL)
		})

		buttonRemove := hb.NewButton().
			Child(hb.I().Class("bi bi-trash")).
			Title("Delete").
			Class("btn btn-sm btn-danger float-end").
			HxInclude("#" + field.form.id).
			HxPost(field.repeaterRemoveUrl + `&repeatable_remove_index=` + cast.ToString(index)).
			HxTarget("#" + field.form.id).
			HxTrigger("click")

		buttonMoveUp := hb.NewButton().
			Child(hb.I().Class("bi bi-arrow-up-circle")).
			Title("Move Up").
			Class("btn btn-sm btn-default").
			HxInclude("#" + field.form.id).
			HxPost(field.repeaterMoveUpUrl + `&repeatable_move_up_index=` + cast.ToString(index)).
			HxTarget("#" + field.form.id).
			HxTrigger("click")

		buttonMoveDown := hb.NewButton().
			Child(hb.I().Class("bi bi-arrow-down-circle")).
			Title("Move Down").
			Class("btn btn-sm btn-default").
			HxInclude("#" + field.form.id).
			HxPost(field.repeaterMoveDownUrl + `&repeatable_move_down_index=` + cast.ToString(index)).
			HxTarget("#" + field.form.id).
			HxTrigger("click")

		card := hb.NewDiv().
			Class("card w-100 mb-3").
			Child(hb.NewDiv().
				Class("card-header").
				Child(buttonMoveUp).
				Child(buttonMoveDown).
				Child(buttonRemove)).
			Child(hb.NewDiv().
				Class("card-body").
				Children(children))

		cards.Child(card)
	}

	return hb.NewDiv().
		Class("form-group mb-3").
		Child(formGroupLabel).
		Child(cards)
}
