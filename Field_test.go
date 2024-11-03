package form

import (
	"strings"
	"testing"
)

func TestFieldBlockEditor(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_BLOCKEDITOR,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label>`,
		`<textarea class="form-control" id="ID" name="NAME">VALUE</textarea>`,
		`<div class="alert alert-danger">CustomInput is nil</div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldDate(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_DATE,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="date" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldDateTime(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_DATETIME,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="datetime-local" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldHidden(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_HIDDEN,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><input class="form-control" id="ID" name="NAME" type="hidden" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldHtmlArea(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_HTMLAREA,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`function initWysiwyg(textareaID, config) {`,
		`<textarea class="form-control" id="ID" name="NAME">VALUE</textarea>`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldImage(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_IMAGE,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><div class="row g-3" style="border: 1px solid silver;border-radius: 10px; margin-top: 0px; margin-left: 0px;margin-right: 0px;"><div class="col-md-2"><img class="img-fluid rounded-start" src="VALUE" style="margin-bottom: 15px;width:100%;max-height:100px;" /></div><div class="col-md-10"><textarea  class="form-control" id="ID" name="NAME" style="height:70px;" type="text">VALUE</textarea><span>The URL can be base64 encoded image URL</span></div></div></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldNumber(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_NUMBER,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="number" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldPassword(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_PASSWORD,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="password" value="VALUE" /></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldRaw(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE VALUE1 <br /> VALUE2 VALUE3",
		Type:  FORM_FIELD_TYPE_RAW,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`VALUE`,
		`VALUE1`,
		`<br />`,
		`VALUE2`,
		`VALUE3`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldSelect(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_SELECT,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><select class="form-select" id="ID" name="NAME"></select></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldSelectWithOptions(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_SELECT,
		Options: []FieldOption{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
			{
				Key:   "VALUE",
				Value: "value3",
			},
		},
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expecteds := []string{
		`<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><select class="form-select" id="ID" name="NAME">`,
		`<option value="key1">value1</option>`,
		`<option value="key2">value2</option>`,
		`<option selected="selected" value="VALUE">value3</option>`,
		`</select></div>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFieldString(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_STRING,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><input class="form-control" id="ID" name="NAME" type="text" value="VALUE" /></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}

func TestFieldTextArea(t *testing.T) {
	field := Field{
		ID:    "ID",
		Name:  "NAME",
		Value: "VALUE",
		Type:  FORM_FIELD_TYPE_TEXTAREA,
	}

	formGroup := field.BuildFormGroup("")

	html := formGroup.ToHTML()

	expected := `<div class="form-group mb-3"><label class="form-label" for="ID">NAME</label><textarea class="form-control" id="ID" name="NAME">VALUE</textarea></div>`
	if html != expected {
		t.Fatal(`Expected: `, expected, ` but was: `, html)
	}
}
