package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MyCtx struct {
	url    *url.URL
	header http.Header
	client *http.Client
	c      context.Context
}

type langKey string
type authKey string

const lang langKey = "lang"
const auth authKey = "auth"

func (m *MyCtx) SetAuthKey(s string) {
	arrs := strings.Split(s,":")
	if len(arrs)!=2 {
		fmt.Println("認証キーの形式が違います。")
		return
	}
	if arrs[1] != "fx"  {
		fmt.Println("認証キーの形式が違います。")
		return
	}
	m.c = context.WithValue(m.c, auth, s)
}

func (m *MyCtx) GetAuthKey() (string, error) {
	v, ok := m.c.Value(auth).(string)
	if !ok {
		return "", fmt.Errorf("failed to get authkey")
	}

	return v, nil
}

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
