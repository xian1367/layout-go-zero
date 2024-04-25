package helper

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

// StrPlural 转为复数 user -> users
func StrPlural(word string) string {
	return inflection.Plural(word)
}

// StrSingular 转为单数 users -> user
func StrSingular(word string) string {
	return inflection.Singular(word)
}

// StrSnake 转为 snake_case，如 TopicComment -> topic_comment
func StrSnake(s string) string {
	return strcase.ToSnake(s)
}

// StrCamel 转为 CamelCase，如 topic_comment -> TopicComment
func StrCamel(s string) string {
	return strcase.ToCamel(s)
}

// StrLowerCamel 转为 lowerCamelCase，如 TopicComment -> topicComment
func StrLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
