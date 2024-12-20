package form

import (
	"strings"
	"testing"
)

func TestFieldRepeater(t *testing.T) {
	repeater := NewRepeater(RepeaterOptions{
		Name: "REPEATER_NAME",
		Fields: []FieldInterface{
			&Field{
				ID:   "ID_1",
				Name: "FIELD_NAME_1",
				Type: FORM_FIELD_TYPE_STRING,
			},
			&Field{
				ID:   "ID_2",
				Name: "FIELD_NAME_2",
				Type: FORM_FIELD_TYPE_STRING,
			},
		},
		Values: []map[string]string{
			{
				"FIELD_NAME_1": "VALUE_1_01",
				"FIELD_NAME_2": "VALUE_2_01",
			},
			{
				"FIELD_NAME_1": "VALUE_1_02",
				"FIELD_NAME_2": "VALUE_2_02",
			},
		},
		RepeaterAddUrl:      "REPEATER_ADD_URL",
		RepeaterMoveUpUrl:   "REPEATER_MOVE_UP_URL",
		RepeaterMoveDownUrl: "REPEATER_MOVE_DOWN_URL",
		RepeaterRemoveUrl:   "REPEATER_REMOVE_URL",
	})

	result := repeater.BuildFormGroup("").ToHTML()

	if result == "" {
		t.Error("Repeater did not return the expected values. Got:", result)
	}

	expecteds := []string{
		`<div class="form-group mb-3">`,
		`<input class="form-control" id="ID_1_0" name="REPEATER_NAME[FIELD_NAME_1][]" type="text" value="VALUE_1_02" />`,
		`<input class="form-control" id="ID_2_1" name="REPEATER_NAME[FIELD_NAME_2][]" type="text" value="VALUE_2_02" />`,
	}

	for _, expected := range expecteds {
		if !strings.Contains(result, expected) {
			t.Fatal(`Expected: `, expected, ` but was: `, result)
		}
	}
}
