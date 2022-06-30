package utils

import (
	"fmt"
	"reflect"
	"strings"
)

var languages = []string{
	"BG",
	"CS",
	"DA",
	"DE",
	"EL",
	"EN",
	"ES",
	"ET",
	"FI",
	"FR",
	"HU",
	"ID",
	"IT",
	"JA",
	"LT",
	"LV",
	"NL",
	"PL",
	"PT",
	"RO",
	"RU",
	"SK",
	"SL",
	"SV",
	"TR",
	"ZH",
}

func ValidateLanguage(s string) {
	s = strings.ToUpper(s)
	switch s {
	case "ENGLISH":
		s = "EN"
	case "日本語":
		s = "JA"
	default:

	}
	if !Contains(languages, s) {
		fmt.Println("その言語は存在しません")
		return
	}
	ctx.SetLanguage(s)
}

func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			// 型変換可能か確認する
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			// 型変換する
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			// 等価判定をする
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}
