package form

import "github.com/gouniverse/hb"

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
	fieldHelp  string
	fieldID    string
	fieldLabel string
	fieldType  string
	fieldName  string
	fieldValue string
	fields     []FieldInterface
	values     [][]map[string]string
}

// == INTERFACE ===============================================================

var _ FieldInterface = (*fieldRepeater)(nil)

// == IMPLEMENTATION OF FieldInterface ========================================
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
	return nil
}
