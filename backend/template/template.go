package template

import (
	"strings"
	"text/template"
	"time"
)

type Template struct {
	*template.Template
}

var Funcs template.FuncMap = make(template.FuncMap)

func init() {
	//register utils func to template
	Funcs["Now"] = func() time.Time {
		return time.Now()
	}
	Funcs["Second"] = func(n int) time.Duration {
		return time.Second * time.Duration(n)
	}
}

func Parse(tmpl string) (*Template, error) {
	tpl, err := template.New("").Funcs(Funcs).Parse(tmpl)
	if err != nil {
		return nil, err
	}
	return &Template{tpl}, nil
}

func (t *Template) ExecuteString() (string, error) {
	var sb strings.Builder
	err := t.Execute(&sb, nil)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
