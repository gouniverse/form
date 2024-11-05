package form

import "github.com/gouniverse/hb"

type FieldInterface interface {
	GetLabel() string
	SetLabel(fieldLabel string)
	GetName() string
	SetName(fieldName string)
	GetRequired() bool
	SetRequired(fieldRequired bool)
	GetType() string
	SetType(fieldType string)
	GetValue() string
	SetValue(fieldValue string)
	BuildFormGroup(fileManagerURL string) *hb.Tag
}
