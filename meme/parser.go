package meme

import (
	"bytes"
	"cryptolist/common"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"html/template"
)

var (
	tmplFuncMap = template.FuncMap{
		"FormatNumber": formatNumber,
	}
)

func ParseCaption(tmpl string, market *common.Market) (string, error) {
	t, err := template.New("meme").Funcs(tmplFuncMap).Parse(tmpl)
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

func formatNumber(f float64) string {
	switch {
	case f > 10000:
		f = f/1000
		return fmt.Sprintf("%.1fK", f)
	case f > 100:
		p := message.NewPrinter(language.English)
		return p.Sprintf("%d", int(f))
	default:
		return fmt.Sprintf("%.2f", f)
	}
}