package main

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Render_tables(info DBInfo, tables []map[string]string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<ul>\n\t")
	for _, table := range tables {

		_buffer.WriteString("<li>\n\t\t\t<a href=\"/dbs/")
		_buffer.WriteString(gorazor.HTMLEscape(info.ID))
		_buffer.WriteString("/tables/")
		_buffer.WriteString(gorazor.HTMLEscape(table["name"]))
		_buffer.WriteString("\">")
		_buffer.WriteString(gorazor.HTMLEscape(table["name"]))
		_buffer.WriteString("</a>\n\t\t</li>")

	}
	_buffer.WriteString("\n</ul>")

	return _buffer.String()
}
