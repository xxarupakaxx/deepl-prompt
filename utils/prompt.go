package utils

import (
	"context"
	"fmt"
	"github.com/c-bata/go-prompt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Header に DeepL-Auth-Key を追加してPOST
// split_sentencesは0に設定するべきかも　"nonewlines" -もあり

type MyPrefix struct {
	livePrefix string
	isEnable   bool
}

func (m *MyPrefix) LivePrefix() (string, bool) {
	return m.livePrefix, m.isEnable
}

var prefix *MyPrefix
var ctx *MyCtx

func NewPrompt() *prompt.Prompt {
	u, _ := url.Parse("https://api-free.deepl.com/v2/")
	ctx = &MyCtx{
		url:    u,
		header: http.Header{},
		client: &http.Client{},
		c:      context.Background(),
	}
	prefix = &MyPrefix{
		livePrefix: "Deepl-prompt >>>",
		isEnable:   true,
	}

	ctx.SetLanguage("EN")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("Welcome to deepl-prompt"),
		prompt.OptionHistory([]string{
			"change EN",
			"translate こんにちは",
			"set [key]",
		}),
		prompt.OptionLivePrefix(prefix.LivePrefix),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.White),
		prompt.OptionSuggestionTextColor(prompt.DarkGray),
		prompt.OptionDescriptionTextColor(prompt.DarkGreen),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
	)

	return p

}

func completer(d prompt.Document) []prompt.Suggest {
	args := strings.Split(d.TextBeforeCursor(), " ")
	//w := d.GetWordBeforeCursor()
	return argumentsCompleter(args)
}

func executor(s string) {
	s = strings.TrimSpace(s)
	cmds := strings.Split(s, " ")
	switch cmds[0] {
	case "exit":
		fmt.Println("おおおおおおおおおおおおおお！")
		os.Exit(0)
	case "translate":
		if len(cmds) <= 1 {
			fmt.Println("翻訳する文がありません")
			return
		}
		key, err := ctx.GetAuthKey()
		if err != nil {
			fmt.Println("認証キーが取得できません。setコマンドで認証キーを設定してください。")
			return
		}

		ctx.header.Set("Authorization", "DeepL-Auth-Key "+key)
		text, err := PostTranslate(cmds[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(text)
	case "change":
		if len(cmds) <= 1 {
			fmt.Println("言語を指定してください")
			return
		}
		ValidateLanguage(cmds[1])
		prefix.livePrefix = "Deepl-prompt To " + cmds[1] + " >>>"
	case "help":
		fmt.Println(HelpMessages())
	case "set":
		ctx.SetAuthKey(cmds[1])
		fmt.Println("認証キーがセットされました")
	}
}
