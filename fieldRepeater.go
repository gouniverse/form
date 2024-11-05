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
func (field *fieldRepeater) GetLabel() string {
	return field.fieldLabel
}

func (field *fieldRepeater) SetLabel(fieldLabel string) {
	field.fieldLabel = fieldLabel
}

func (field *fieldRepeater) GetName() string {
	return field.fieldName
}

func (field *fieldRepeater) SetName(fieldName string) {
	field.fieldName = fieldName
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
