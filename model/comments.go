package model

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func (md *MetaData) renderComment(s string) string {

	t := strings.TrimSpace(s)

	switch {
	case t == "":
		return ""
	case md.Cfg.CommentsFormat == "markdown":
		unsafe := blackfriday.Run([]byte(t))
		html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

		html = strings.TrimRight(html, "\r\n")
		t = strings.TrimRight(t, "\r\n")

		if !strings.ContainsAny(t, "\r\n") {
			html = strings.TrimPrefix(html, "<p>")
			html = strings.TrimSuffix(html, "</p>")
		}

		return html
	default:
		return string(bluemonday.UGCPolicy().SanitizeBytes([]byte(t)))
	}
	return ""
}
