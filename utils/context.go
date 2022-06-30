package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type MyCtx struct {
	url    *url.URL
	header http.Header
	client *http.Client
	c      context.Context
}

type langKey string

const lang langKey = "lang"

func (m *MyCtx) SetLanguage(s string) {
	m.c = context.WithValue(m.c, lang, s)
}

func (m *MyCtx) GetLanguage() (string, error) {
	v, ok := m.c.Value(lang).(string)
	if !ok {
		return "", fmt.Errorf("failed to get language")
	}

	return v, nil
}
