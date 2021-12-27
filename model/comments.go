package model

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func (md *MetaData) renderComment(s string) string {

	switch {
	case s == "", strings.TrimSpace(s) == "":
		return ""
	case md.Cfg.CommentsFormat == "markdown":
		unsafe := blackfriday.Run([]byte(s))
		html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

        if ! strings.ContainsAny(s, "\r\n") {
            html = strings.TrimLeft(html, "<p>")
            html = strings.TrimRight(html, "</p>\n")
        }

		return html
    default:
		return string(bluemonday.UGCPolicy().SanitizeBytes([]byte(s)))
	}
	return ""
}
