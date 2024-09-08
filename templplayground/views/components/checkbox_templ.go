// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
// Code generated by "templplaygroundgenerator"; DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/jfbus/templ-components/components/checkbox"
	"github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/selectfield"
	"github.com/jfbus/templ-components/components/selectfield/option"
	"github.com/jfbus/templ-components/components/size"
	"strconv"
)

func CheckboxViewer(def checkbox.D) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = checkbox.C(def).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CheckboxForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = input.C(input.D{
			Name:  "Name",
			Label: "Name",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = selectfield.C(selectfield.D{
			Name:  "Style",
			Label: "Style",
			Options: []option.D{
				{
					Label: "Select a value",
				},
				{
					Label: "checkbox.StyleBordered",
					Value: strconv.Itoa(int(checkbox.StyleBordered)),
				},
				{
					Label: "checkbox.StyleGrouped",
					Value: strconv.Itoa(int(checkbox.StyleGrouped)),
				},
				{
					Label: "checkbox.StyleGroupedVertical",
					Value: strconv.Itoa(int(checkbox.StyleGroupedVertical)),
				},
				{
					Label: "checkbox.StyleLabelOnly",
					Value: strconv.Itoa(int(checkbox.StyleLabelOnly)),
				},
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = input.C(input.D{
			Name:  "Label",
			Label: "Label",
			Value: "Label",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = input.C(input.D{
			Name:  "Value",
			Label: "Value",
			Value: "Value",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = checkbox.C(checkbox.D{
			Name:  "Checked",
			Label: "Checked",
			Value: "true",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = checkbox.C(checkbox.D{
			Name:  "Disabled",
			Label: "Disabled",
			Value: "true",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = selectfield.C(selectfield.D{
			Name:  "Size",
			Label: "Size",
			Options: []option.D{
				{
					Label: "Select a value",
				},
				{
					Label: "size.Inherit",
					Value: strconv.Itoa(int(size.Inherit)),
				},
				{
					Label: "size.XS",
					Value: strconv.Itoa(int(size.XS)),
				},
				{
					Label: "size.S",
					Value: strconv.Itoa(int(size.S)),
				},
				{
					Label: "size.Normal",
					Value: strconv.Itoa(int(size.Normal)),
				},
				{
					Label: "size.L",
					Value: strconv.Itoa(int(size.L)),
				},
				{
					Label: "size.XL",
					Value: strconv.Itoa(int(size.XL)),
				},
				{
					Label: "size.TwoXL",
					Value: strconv.Itoa(int(size.TwoXL)),
				},
				{
					Label: "size.ThreeXL",
					Value: strconv.Itoa(int(size.ThreeXL)),
				},
				{
					Label: "size.FourXL",
					Value: strconv.Itoa(int(size.FourXL)),
				},
				{
					Label: "size.FiveXL",
					Value: strconv.Itoa(int(size.FiveXL)),
				},
				{
					Label: "size.SixXL",
					Value: strconv.Itoa(int(size.SixXL)),
				},
				{
					Label: "size.SevenXL",
					Value: strconv.Itoa(int(size.SevenXL)),
				},
				{
					Label: "size.EightXL",
					Value: strconv.Itoa(int(size.EightXL)),
				},
				{
					Label: "size.NineXL",
					Value: strconv.Itoa(int(size.NineXL)),
				},
				{
					Label: "size.Full",
					Value: strconv.Itoa(int(size.Full)),
				},
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CheckboxSection() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = ComponentViewer(Component{
			ID:   "Checkbox",
			Form: CheckboxForm(),
			Viewer: CheckboxViewer(checkbox.D{
				Label: "Label",
				Value: "Value",
			}),
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
