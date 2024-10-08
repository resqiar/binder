// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "binder/views/components/navbar"
import "binder/views/components/input"

func CreateExtPage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IndexHead("Create new Extension | Binder").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><link href=\"/static/ace/theme/terminal.min.css\" rel=\"stylesheet\"></head><body class=\"p-2 font-bold\"><header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = navbar.CreateNavbar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateExtBody().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main><script src=\"/static/dropzone.min.js\"></script><script src=\"/static/ace/ace.js\"></script><script src=\"/static/editor.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CreateExtBody() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-post=\"/api/ext/create\" hx-target=\"#info-wrapper\" hx-indicator=\"#loading\" hx-disabled-elt=\"button[type=&#39;submit&#39;]\" enctype=\"multipart/form-data\" class=\"my-12 flex w-full flex-col items-center py-4 px-4 lg:px-12\"><h1 class=\"text-2xl font-bold\">Create a New Extension</h1><div class=\"flex w-full flex-col gap-4 py-8 pb-20 lg:w-8/12\"><!-- Title --><div class=\"form-control\"><label class=\"label\" for=\"title-input\"><span class=\"label-text\">Title of Extension*</span></label> <label class=\"input input-bordered flex items-center gap-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M4 6h16M4 12h8m-8 6h16\"></path></svg> <input type=\"text\" name=\"ext-title\" class=\"grow\" placeholder=\"Title\" required></label> <label class=\"label px-2\" for=\"title-input\"><span class=\"label-text-alt\">Title must be at least 3 characters</span></label></div><!-- Description Textarea --><div class=\"form-control\"><label class=\"label\" for=\"desc-input\"><span class=\"label-text\">Description</span></label> <textarea id=\"desc-input\" name=\"ext-desc\" class=\"textarea textarea-bordered w-full\" rows=\"10\" placeholder=\"Description\"></textarea></div><!-- Drag Drop Image --><div class=\"form-control\"><label class=\"label\" for=\"dropzone\"><span class=\"label-text\">Upload Image</span></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = input.ImageDropzone().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Youtube URL --><div class=\"form-control\"><label class=\"label\" for=\"yt-input\"><span class=\"label-text\">Youtube Video</span></label> <label class=\"input input-bordered flex items-center gap-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"h-5 w-5\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M13.19 8.688a4.5 4.5 0 011.242 7.244l-4.5 4.5a4.5 4.5 0 01-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 00-6.364-6.364l-4.5 4.5a4.5 4.5 0 001.242 7.244\"></path></svg> <input type=\"text\" id=\"yt-input\" name=\"ext-yt\" placeholder=\"Youtube video URL\" class=\"grow\"></label></div><!-- CODE INPUT SECTION --><section><label class=\"label\" for=\"editor\"><span class=\"label-text\">Code</span></label> <input aria-hidden=\"true\" type=\"hidden\" name=\"ext-code\" id=\"hidden-editor\"><div id=\"editor\" class=\"min-h-[500px] !text-lg\"></div><label class=\"label\" for=\"editor\"><span class=\"label-text-alt\"><span id=\"editor-len\">0</span>/10000</span></label></section><div id=\"info-wrapper\"></div><div class=\"my-12 w-full flex justify-end\"><button type=\"submit\" class=\"w-full btn btn-primary\">Create <span id=\"loading\" class=\"loading loading-dots loading-md loading-indicator\"></span></button></div></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
