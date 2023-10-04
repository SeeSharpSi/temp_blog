// Code generated by templ@v0.2.364 DO NOT EDIT.

package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"

import types "blog/main/types"

func Posts(temp []types.Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"page\">")
		if err != nil {
			return err
		}
		for _, v := range temp {
			_, err = templBuffer.WriteString("<div class=\"post\" hx-target=\"#page\" hx-swap=\"outerHTML\" hx-get=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString("/post/" + strconv.Itoa(v.Id)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"><div class=\"title\">")
			if err != nil {
				return err
			}
			var var_2 string = v.Title
			_, err = templBuffer.WriteString(templ.EscapeString(var_2))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div class=\"content\">")
			if err != nil {
				return err
			}
			var var_3 string = v.Content
			_, err = templBuffer.WriteString(templ.EscapeString(var_3))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
