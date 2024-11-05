package form

import "github.com/gouniverse/hb"

type FieldInterface interface {
	GetID() string
	SetID(fieldID string)
	GetLabel() string
	SetLabel(fieldLabel string)
	GetName() string
	SetName(fieldName string)
	GetHelp() string
	SetHelp(fieldName string)
	GetOptions() []FieldOption
	SetOptions(fieldOptions []FieldOption)
	GetOptionsF() func() []FieldOption
	SetOptionsF(fieldOptionsF func() []FieldOption)
	GetRequired() bool
	SetRequired(fieldRequired bool)
	GetType() string
	SetType(fieldType string)
	GetValue() string
	SetValue(fieldValue string)
	BuildFormGroup(fileManagerURL string) *hb.Tag
}
