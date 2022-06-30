package utils

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Header に DeepL-Auth-Key を追加してPOST
// split_sentencesは0に設定するべきかも　"nonewlines" -もあり

type MyCtx struct {
	url    *url.URL
	header http.Header
	client *http.Client
}

type MyPrefix struct {
	livePrefix string
	isEnable   bool
}

func (m *MyPrefix) LivePrefix() (string, bool) {
	return m.livePrefix, m.isEnable
}

var livePrefix *MyPrefix
var ctx *MyCtx

func NewPrompt() *prompt.Prompt {
	u, _ := url.Parse("https://api-free.deepl.com/v2/")
	ctx = &MyCtx{
		url:    u,
		header: http.Header{},
		client: &http.Client{},
	}
	livePrefix = &MyPrefix{
		livePrefix: ">>>",
		isEnable:   false,
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionHistory([]string{
			"change EN",
			"translate こんにちは",
		}),
		prompt.OptionDescriptionBGColor(prompt.Brown),
		prompt.OptionPrefix("Deepl-prompt >>>"),
		prompt.OptionLivePrefix(livePrefix.LivePrefix),
		prompt.OptionMaxSuggestion(2),
	)

	return p

}

func completer(d prompt.Document) []prompt.Suggest {
	s := d.GetWordBeforeCursor()
	return prompt.FilterContains(GetSuggestions(), s, true)
}

func executor(s string) {
	s = strings.TrimSpace(s)
	commands := strings.Split(s," ")
	switch commands[0] {
	case "exit":
		fmt.Println("おおおおおおおおおおおおおお！")
		os.Exit(0)
	case "translate":


	}
}
