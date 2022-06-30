package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

type Response struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	Language string `json:"detected_source_language"`
	Text     string `json:"text"`
}

func PostTranslate(s string) (string, error) {
	url, err := url2.Parse(ctx.url.String() + "translate")
	fmt.Println(url)
	if err != nil {

		return "", fmt.Errorf("failed to url parse :%w", err)
	}

	lan, err := ctx.GetLanguage()

	if err != nil {
		return "", fmt.Errorf("failed to get language :%w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to request : %w", err)
	}

	req.Header = ctx.header

	q := req.URL.Query()
	q.Add("target_lang", lan)
	q.Add("text", s)

	req.URL.RawQuery = q.Encode()

	res, err := ctx.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to response :%w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body :%w", err)
	}

	if res.StatusCode >= 400 {

		return "", fmt.Errorf("failed to post : badrequest: %v", string(body))
	}

	return string(body), nil
}
