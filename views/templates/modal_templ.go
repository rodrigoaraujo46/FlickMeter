// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func modal() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<dialog tabindex=\"-1\" class=\"font-mono w-3/4 md:w-[650px] focus:outline-none overflow-hidden rounded-lg shadow-md shadow-black\" id=\"modal\" hx-on::after-settle=\"event.target.id === &#39;modal&#39; ? toggleModal() : null\"></dialog><script>\n    const modal = document.getElementById('modal');\n    modal.addEventListener('keydown', event => {\n        if (!modal.open) {\n            return\n        }\n        if (event.key === 'Escape') {\n            event.preventDefault();\n\n            const isAnimationRunning = !modal.classList.contains('closing') && modal.getAnimations().some(animation => animation.playState === 'running');\n            if (isAnimationRunning) {\n                modal.addEventListener('animationend', () => {toggleModal();}, {once: true});\n                return\n            }\n            toggleModal();\n        }\n\n    });\n    function toggleModal() {\n        if (modal.open) {\n            modal.focus()\n            modal.addEventListener('animationend', () => {\n                modal.close();\n                modal.classList.remove('closing');\n            }, {once: true});\n\n            modal.classList.add('closing');\n            return\n        }\n\n        const input = document.getElementById('form').querySelector('input');\n\n        modal.addEventListener('animationend', function () {\n            input.focus();\n        }, {once: true});\n        modal.showModal();\n        modal.focus();\n    }\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
