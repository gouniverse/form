package form

import (
	"strings"
	"testing"
)

func TestForm(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			&Field{
				ID:    "ID",
				Name:  "NAME",
				Value: "VALUE",
				Type:  FORM_FIELD_TYPE_STRING,
			},
		},
	})

	html := form.Build().ToHTML()

	if html == "" {
		t.Fatal("Expected html to not be nil")
	}

	expecteds := []string{
		`<form method="POST">`,
		`<div class="form-group mb-3">`,
		`<label class="form-label" for="ID">NAME</label>`,
		`<input class="form-control" id="ID" name="NAME" type="text" value="VALUE" />`,
		`</div>`,
		`</form>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}

func TestFormWithRepeater(t *testing.T) {
	form := NewForm(FormOptions{
		Fields: []FieldInterface{
			NewRepeater(RepeaterOptions{
				Label: "LABEL",
				Type:  FORM_FIELD_TYPE_STRING,
				Fields: []FieldInterface{
					&Field{
						ID:    "ID",
						Name:  "NAME",
						Value: "VALUE",
						Type:  FORM_FIELD_TYPE_STRING,
					}},
				Values: [][]map[string]string{},
			}),
		},
	})

	html := form.Build().ToHTML()

	if html == "" {
		t.Fatal("Expected html to not be nil")
	}

	expecteds := []string{
		`<form method="POST">`,
		`</form>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}
