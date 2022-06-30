package utils

import "github.com/c-bata/go-prompt"

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterContains(commands, args[0], true)
	}

	switch args[0] {
	case "change":
		if len(args) == 2 {
			subCommands := []prompt.Suggest{
				{Text: "EN", Description: "英語"},
				{Text: "BG"},
				{Text: "CS"},
				{Text: "DA"},
				{Text: "DE"},
				{Text: "EL"},
				{Text: "ES"},
				{Text: "ET"},
				{Text: "FL"},
				{Text: "FL"},
				{Text: "FR"},
				{Text: "HU"},
				{Text: "ID"},
				{Text: "IT"},
				{Text: "JA"},
				{Text: "LT"},
				{Text: "LV"},
				{Text: "NL"},
				{Text: "PL"},
				{Text: "PT"},
				{Text: "RO"},
				{Text: "RU"},
				{Text: "SK"},
				{Text: "SL"},
				{Text: "SV"},
				{Text: "TR"},
				{Text: "ZH"},
			}
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	}
	return []prompt.Suggest{}
}

var commands = []prompt.Suggest{
	{
		Text:        "change",
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
	{
		Text:        "usage",
		Description: "翻訳量や設定された制限を監視することができます",
	},
}
