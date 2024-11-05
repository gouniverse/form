package form

import "github.com/gouniverse/hb"

type FieldInterface interface {
	GetType() string
	GetName() string
	GetValue() string
	SetType(fieldType string)
	SetName(fieldName string)
	SetValue(fieldValue string)
	BuildFormGroup(fileManagerURL string) *hb.Tag
}
