# Form

<a href="https://gitpod.io/#https://github.com/gouniverse/form" target="_blank" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

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

```
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

```
buttonSave := hb.NewButton().
	Class("btn btn-primary ms-2 float-end").
	Child(hb.NewI().Class("bi bi-save").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
	HTML("Save").
	HxInclude("#FormCustomerUpdate").
	HxPost(links.NewAdminLinks().CustomerUpdate(data.customerID)).
	HxTarget("#FormCustomerUpdate")
```
