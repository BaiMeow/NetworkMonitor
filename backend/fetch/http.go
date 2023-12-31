package fetch

import (
	"github.com/BaiMeow/NetworkMonitor/template"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

func init() {
	Register("http", func(m map[string]any) (Fetcher, error) {
		url, err := template.ParseInterface(m["url"])
		if err != nil {
			return nil, errors.Wrap(err, "field url")
		}

		var body *template.Template
		if m["body"] != nil {
			body, err = template.ParseInterface(m["body"])
			if err != nil {
				return nil, errors.Wrap(err, "field body")
			}
		}

		header := &http.Header{}
		if m["header"] != nil {
			headerMap, ok := m["header"].(map[string]any)
			if !ok {
				return nil, errors.New("field header should be map[string](string|[]string)")
			}
			for k, v := range headerMap {
				switch v := v.(type) {
				case string:
					header.Add(k, v)
				case []string:
					for _, sv := range v {
						header.Add(k, sv)
					}
				default:
					return nil, errors.New("field header should be map[string](string|[]string)")
				}
			}
		}
		method, ok := m["method"].(string)
		if !ok {
			return nil, errors.New("method should be string")
		}
		h := &HTTP{
			URL:    url,
			Body:   body,
			Header: header,
			Method: method,
		}
		return h, nil
	})
}

var _ Fetcher = (*HTTP)(nil)

type HTTP struct {
	URL    *template.Template
	Body   *template.Template
	Header *http.Header
	Method string
}

func (f *HTTP) GetData() ([]byte, error) {
	url, err := f.URL.ExecuteString()
	if err != nil {
		return nil, errors.Wrap(err, "generate url")
	}

	var body io.Reader
	if f.Body != nil {
		bodyStr, err := f.Body.ExecuteString()
		if err != nil {
			return nil, errors.Wrap(err, "generate body")
		}
		bodyReader := strings.NewReader(bodyStr)
		body = bodyReader
	}

	req, err := http.NewRequest(f.Method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "fail to new http.Request")
	}

	if f.Header != nil {
		req.Header = f.Header.Clone()
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "fail to do http")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read http body")
	}

	return data, nil
}
