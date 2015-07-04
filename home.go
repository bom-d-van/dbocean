package main

import (
	"bytes"
)

func Home(body string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n\t<title>DB Ocean</title>\n</head>\n<body>\n\t<h1><a href=\"/dbs\">DB Ocean</a></h1>\n\t")
	_buffer.WriteString((body))
	_buffer.WriteString("\n</body>\n</html>")

	return _buffer.String()
}
