package editor

import (
	"text/template"
	"testing"
)

func TestBasicHtmlProcessing(t *testing.T) {
  var html string
  tmpl := template.Must(template.New("test").Parse("{{.Name}}:{{.Body}}"))
  html = processHtml("foobar", tmpl)
  expected := "test:foobar"

  if html != expected {
    t.Error("Expected", expected, "got", html)
  }
}
