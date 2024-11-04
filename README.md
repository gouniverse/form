# Form

<a href="https://gitpod.io/#https://github.com/gouniverse/form" target="_blank" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

## License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0). You can find a copy of the license at [https://www.gnu.org/licenses/gpl-3.0.en.html](https://www.gnu.org/licenses/agpl-3.0.txt)

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Form Field Types

- date
- datetime
- image
- htmlarea
- hidden
- number
- password
- select
- string
- table
- textarea
- raw

## Example Customer Update Form

This is an example taken from real life code of a controller type with a form method. The method returns a customer update Form.

```golang
func (controller customerUpdateController) form(data customerUpdateControllerData) *hb.Tag {
	updateCustomerForm := form.NewForm(form.FormOptions{
		ID: "FormCustomerUpdate",
	})
	updateCustomerForm.SetFields([]form.Field{
		{
			Label: "Status",
			Name:  "customer_status",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Value: data.formStatus,
			Options: []form.FieldOption{
				{
					Value: "Draft",
					Key:   models.CUSTOMER_STATUS_DRAFT,
				},
				{
					Value: "Active",
					Key:   models.CUSTOMER_STATUS_ACTIVE,
				},
				{
					Value: "Inactive",
					Key:   models.CUSTOMER_STATUS_INACTIVE,
				},
				{
					Value: "Deleted",
					Key:   models.CUSTOMER_STATUS_DELETED,
				},
			},
		},
		{
			Label: "Type",
			Name:  "customer_type",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formType,
		},
		{
			Label: "First Name",
			Name:  "first_name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formPersonFirstName,
		},
		{
			Label: "Last Name",
			Name:  "last_name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formPersonLastName,
		},
		{
			Label: "Company Name",
			Name:  "company_name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formCompanyName,
		},
		{
			Label: "Country",
			Name:  "country",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Value: data.formCountry,
			Options: lo.Map(data.countries, func(country models.Country, _ int) form.FieldOption {
				return form.FieldOption{
					Key:   country.IsoCode2(),
					Value: country.Name(),
				}
			}),
		},
		{
			Label:    "Customer ID",
			Name:     "customer_id",
			Type:     form.FORM_FIELD_TYPE_STRING,
			Value:    data.customerID,
			Readonly: true,
		},
	})

	if data.formErrorMessage != "" {
		updateCustomerForm.AddField(form.Field{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: swal.Swal(swal.SwalOptions{Icon: "error", Text: data.formErrorMessage}),
		})
	}

	if data.formSuccessMessage != "" {
		updateCustomerForm.AddField(form.Field{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: swal.Swal(swal.SwalOptions{Icon: "success", Text: data.formSuccessMessage}),
		})
	}

	return updateCustomerForm.Build()

}
```

The save button outside of the form method. It uses HTMX for submitting the form.

```golang
buttonSave := hb.NewButton().
	Class("btn btn-primary ms-2 float-end").
	Child(hb.NewI().Class("bi bi-save").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
	HTML("Save").
	HxInclude("#FormCustomerUpdate").
	HxPost(links.NewAdminLinks().CustomerUpdate(data.customerID)).
	HxTarget("#FormCustomerUpdate")
```

## Advanced

### Configuring the HTML textarea. 

The HTML area uses the Trumbowyg wisiwyg editor. To use it you need to add the following to the web page:

```html
// Trumbowyg requires jQuery >= 1.8
<script src="//ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="js/vendor/jquery-3.3.1.min.js"><\/script>')</script>
<script src="trumbowyg/dist/trumbowyg.min.js"></script>
<link rel="stylesheet" href="trumbowyg/dist/ui/trumbowyg.min.css">
```

If you need special options
for the editor, you can use the `Options` field as in the example below.

```golang
{
	Label: "Content",
	Name:  "post_content",
	Type:  FORM_FIELD_TYPE_TEXTAREA,
	Value: data.formContent,
	Help:  "The content of this blog post to display on the post details page.",
	Options: []form.FieldOption{
		{
			Key: "config",
			Value: `{
btns: [
	['viewHTML'],
	['undo', 'redo'],
	['formatting'],
	['strong', 'em', 'del'],
	['superscript', 'subscript'],
	['link','justifyLeft','justifyRight','justifyCenter','justifyFull'],
	['unorderedList', 'orderedList'],
	['insertImage'],
	['removeformat'],
	['horizontalRule'],
	['fullscreen'],
],
autogrow: true,
removeformatPasted: true,
tagsToRemove: ['script', 'link', 'embed', 'iframe', 'input'],
tagsToKeep: ['hr', 'img', 'i'],
autogrowOnEnter: true,
linkTargets: ['_blank'],
}`,
		},
},
```
