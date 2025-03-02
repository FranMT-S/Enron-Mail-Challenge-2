package helpers

import "fmt"

var ValueInnerParentheses = `\([^)]+\)`
var NameFieldWithSpecialCharacters = `[-*]?\w+`
var FieldRegex = fmt.Sprintf(`%v:(%v|\S+)`, NameFieldWithSpecialCharacters, ValueInnerParentheses)
