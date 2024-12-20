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
				Name:                "REPEATER_NAME",
				Label:               "LABEL",
				RepeaterAddUrl:      "REPEATER_ADD_URL",
				RepeaterMoveUpUrl:   "REPEATER_MOVE_UP_URL",
				RepeaterMoveDownUrl: "REPEATER_MOVE_DOWN_URL",
				RepeaterRemoveUrl:   "REPEATER_REMOVE_URL",
				Fields: []FieldInterface{
					&Field{
						ID:    "ID_1",
						Name:  "NAME_1",
						Value: "VALUE_1",
						Type:  FORM_FIELD_TYPE_STRING,
					},
					&Field{
						ID:    "ID_2",
						Name:  "NAME_2",
						Value: "VALUE_2",
						Type:  FORM_FIELD_TYPE_STRING,
					},
				},
				Values: []map[string]string{
					{
						"NAME_1": "VALUE_1_01",
						"NAME_2": "VALUE_2_01",
					},
					{
						"NAME_1": "VALUE_1_02",
						"NAME_2": "VALUE_2_02",
					},
				},
			}),
		},
	})

	html := form.Build().ToHTML()

	if html == "" {
		t.Fatal("Expected html to not be nil")
	}

	expecteds := []string{
		`<form method="POST">`,
		`<div class="form-group mb-3">`,
		`<label class="form-label">LABEL`,
		`hx-post="REPEATER_ADD_URL"`,
		`hx-post="REPEATER_REMOVE_URL`,
		`hx-post="REPEATER_MOVE_UP_URL`,
		`hx-post="REPEATER_MOVE_DOWN_URL`,
		`name="REPEATER_NAME[NAME_1][]"`,
		`name="REPEATER_NAME[NAME_2][]"`,
		`value="VALUE_1_01"`,
		`value="VALUE_1_02"`,
		`value="VALUE_2_01"`,
		`value="VALUE_2_02"`,
		`</form>`,
	}
	for _, expected := range expecteds {
		if !strings.Contains(html, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, html)
		}
	}
}
