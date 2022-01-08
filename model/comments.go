package model

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// renderComment checks, and renders, a comment into the appropriate html.
// Current rendering options are "none" and "markdown"
func (md *MetaData) renderComment(s string) string {

	t := strings.TrimSpace(s)

	switch {
	case t == "":
		return ""
	case md.Cfg.CommentsFormat == "markdown":
		unsafe := blackfriday.Run([]byte(t))
		html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
		html = strings.TrimSpace(html)

		// If the comment resulted in a single html line, no more, then remove
		// the wrapping para tags
		trimMl := strings.TrimPrefix(html, "<p>")
		trimMl = strings.TrimSuffix(trimMl, "</p>")

		if strings.Count(trimMl, "<p>") == 0 {
			return trimMl
		}

		return html
	default:
		return string(bluemonday.UGCPolicy().SanitizeBytes([]byte(t)))
	}
	return ""
}
