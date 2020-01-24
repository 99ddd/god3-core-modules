package language

import (
	"github.com/99ddd/god3-core-config/config"
	"golang.org/x/text/language"
	"html/template"
	"strings"
)

// Ref:
// https://github.com/GoAdminGroup/go-admin/blob/master/modules/language

var (
	EN = language.English.String()
)

// Get return the value of default scope.
func Get(value string) string {
	return GetWithScope(value)
}

// GetWithScope return the value of given scopes.
func GetWithScope(value string, scopes ...string) string {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// GetFromHtml return the value of given scopes and template.HTML value.
func GetFromHtml(value template.HTML, scopes ...string) template.HTML {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := Lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(string(value))]; ok {
		return template.HTML(locale)
	}

	return value
}

// WithScopes join scopes prefix and the value.
func WithScopes(value string, scopes ...string) string {
	return JoinScopes(scopes) + strings.ToLower(value)
}

// LangMap is the map of language packages.
type LangMap map[string]map[string]string

// Lang is the global LangMap.
var Lang = LangMap{
	language.English.String(): en,

	"en": en,
}

// Get get the value from LangMap.
func (lang LangMap) Get(value string) string {
	return lang.GetWithScope(value)
}

// GetWithScope get the value from LangMap with given scopes.
func (lang LangMap) GetWithScope(value string, scopes ...string) string {
	if config.Get().Language == "" {
		return value
	}

	if locale, ok := lang[config.Get().Language][JoinScopes(scopes)+strings.ToLower(value)]; ok {
		return locale
	}

	return value
}

// Add add a language package to the Lang.
func Add(key string, lang map[string]string) {
	Lang[key] = lang
}

func JoinScopes(scopes []string) string {
	j := ""
	for _, scope := range scopes {
		j += scope + "."
	}
	return j
}
