package form

func NewRepeater(opts RepeaterOptions) *fieldRepeater {
	return &fieldRepeater{
		fieldHelp:           opts.Help,
		fieldLabel:          opts.Label,
		fieldName:           opts.Name,
		fieldType:           form_FIELD_TYPE_REPEATER,
		fieldValue:          opts.Value,
		fields:              opts.Fields,
		values:              opts.Values,
		repeaterAddUrl:      opts.RepeaterAddUrl,
		repeaterMoveUpUrl:   opts.RepeaterMoveUpUrl,
		repeaterMoveDownUrl: opts.RepeaterMoveDownUrl,
		repeaterRemoveUrl:   opts.RepeaterRemoveUrl,
	}
}

type RepeaterOptions struct {
	Label               string
	Type                string
	Name                string
	Value               string
	Help                string
	Fields              []FieldInterface
	Values              []map[string]string
	RepeaterAddUrl      string
	RepeaterMoveUpUrl   string
	RepeaterMoveDownUrl string
	RepeaterRemoveUrl   string
}
