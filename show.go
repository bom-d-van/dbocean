package main

import (
	"bytes"
)

func Show(info DBInfo, tables []map[string]string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString((Render_tables(info, tables)))

	return _buffer.String()
}
