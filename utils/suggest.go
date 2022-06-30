package utils

import "github.com/c-bata/go-prompt"

func GetSuggestions() []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{
			Text:        "Change",
			Description: "翻訳する言語を変更します。",
		},
		{
			Text:        "translate",
			Description: "翻訳します",
		},
		{
			Text:        "exit",
			Description: "deepl-promptからばいばい",
		},
	}
	return suggestions
}
