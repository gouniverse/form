package form

import "github.com/gouniverse/hb"

func NewField(opts FieldOptions) *Field {
	return &Field{
		ID:           opts.ID,
		Type:         opts.Type,
		Name:         opts.Name,
		Label:        opts.Label,
		Help:         opts.Help,
		Options:      opts.Options,
		OptionsF:     opts.OptionsF,
		Value:        opts.Value,
		Required:     opts.Required,
		Readonly:     opts.Readonly,
		Disabled:     opts.Disabled,
		TableOptions: opts.TableOptions,
		Placeholder:  opts.Placeholder,
		Invisible:    opts.Invisible,
		CustomInput:  opts.CustomInput,
	}
}

type FieldOptions struct {
	ID           string // automatic, if not assigned
	Type         string
	Name         string
	Label        string
	Help         string
	Options      []FieldOption
	OptionsF     func() []FieldOption
	Value        string
	Required     bool
	Readonly     bool
	Disabled     bool
	TableOptions TableOptions
	Placeholder  string
	Invisible    bool
	CustomInput  hb.TagInterface
}
