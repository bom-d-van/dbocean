package main

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Index(infos []DBInfo) string {
	var _buffer bytes.Buffer
	for _, info := range infos {

		_buffer.WriteString("<a href=\"/dbs/")
		_buffer.WriteString(gorazor.HTMLEscape(info.ID))
		_buffer.WriteString("\">")
		_buffer.WriteString(gorazor.HTMLEscape(info.Name))
		_buffer.WriteString("</a>")

	}
	_buffer.WriteString("\n\n<form method=\"POST\" action=\"/dbs\">\n\t<p>create a new database</p>\n\t<div>\n\t\t<textarea></textarea>\n\t</div>\n\t<button>save</button>\n</form>")

	return _buffer.String()
}
