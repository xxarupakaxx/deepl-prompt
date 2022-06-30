package utils

import (
	"context"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/joho/godotenv"
	"log"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx.header.Set("Authorization", "DeepL-Auth-Key "+os.Getenv("auth_key"))
	ctx.SetLanguage("EN-US")
	p := prompt.New(
		executor,
		completer,
		prompt.OptionHistory([]string{
			"change EN",
			"translate こんにちは",
		}),
		prompt.OptionLivePrefix(prefix.LivePrefix),
		prompt.OptionInputTextColor(prompt.Yellow),
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
		text, err := PostTranslate(cmds[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(text)
	case "change":
		ChangeLanguage(cmds[1])
		prefix.livePrefix = "Deepl-prompt To " + cmds[1] + " >>>"
	}
}
