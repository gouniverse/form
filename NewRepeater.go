package form

func NewRepeater(opts RepeaterOptions) *fieldRepeater {
	return &fieldRepeater{
		fieldType:  opts.Type,
		fieldName:  opts.Name,
		fieldValue: opts.Value,
		fields:     opts.Fields,
		values:     opts.Values,
	}
}

type RepeaterOptions struct {
	Label  string
	Type   string
	Name   string
	Value  string
	Fields []FieldInterface
	Values [][]map[string]string
}
