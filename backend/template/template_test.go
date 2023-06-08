package template

import (
	"bytes"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	var buf bytes.Buffer
	tmpl, err := template.New("").Funcs(Funcs).Parse("{{ (Now.Add (Second -60)).Format \"2006-01-02 15:04\"}}")
	if err != nil {
		t.Error(err)
		return
	}
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(buf.String())
}
