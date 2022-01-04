package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	html := `{"Name": "<b>HTML content</b>"}`
	json.HTMLEscape(&out, []byte(html))
	out.WriteTo(os.Stdout)
}
