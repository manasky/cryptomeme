package meme

import (
	"bytes"
	"cryptolist/common"
	"html/template"
)

func ParseCaption(tmpl string, market *common.Market) (string, error) {
	t, err := template.New("meme").Parse(tmpl)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, market)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
