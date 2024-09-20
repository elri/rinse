package rinser

import (
	"fmt"
	"html"
	"html/template"
	"path/filepath"

	"github.com/linkdata/jaws"
)

type uiJobLink struct{ *Job }

// JawsGetHtml implements jaws.HtmlGetter.
func (ui uiJobLink) JawsGetHtml(rq *jaws.Element) template.HTML {
	var s string
	if ui.State() == JobFinished {
		s = fmt.Sprintf(`<a target="_blank" href="/job/%s">%s</a>`, ui.UUID, html.EscapeString(ui.ResultName()))
	} else {
		s = html.EscapeString(ui.Name)
	}
	s += fmt.Sprintf(`<span class="ms-2 badge text-bg-light">%s</span><span class="ms-2 badge text-bg-light">%s</span>`,
		filepath.Ext(ui.DocumentName()), ui.LanguageName(ui.Lang()))
	return template.HTML(s) // #nosec G203
}

func (job *Job) UiLink() jaws.HtmlGetter {
	return uiJobLink{job}
}