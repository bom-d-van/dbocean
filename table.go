package main

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Table(info DBInfo, tables []map[string]string, table []map[string]interface{}, data []map[string]interface{}) string {
	var _buffer bytes.Buffer
	_buffer.WriteString((Render_tables(info, tables)))
	_buffer.WriteString("\n\n<table>\n\t<thead>\n\t\t<tr>\n\t\t\t")
	for _, column := range table {

		_buffer.WriteString("<th>")
		_buffer.WriteString(gorazor.HTMLEscape(column["name"]))
		_buffer.WriteString("</th>")

	}
	_buffer.WriteString("\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t")
	for _, row := range data {

		_buffer.WriteString("<tr>\n\t\t\t\t")
		for _, column := range row {

			_buffer.WriteString("<td>")
			_buffer.WriteString(gorazor.HTMLEscape(column))
			_buffer.WriteString("</td>")

		}
		_buffer.WriteString("\n\t\t\t</tr>")

	}
	_buffer.WriteString("\n\t</tbody>\n</table>")

	return _buffer.String()
}
