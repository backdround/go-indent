// indent implements Indent function that allows to indent multiline strings
package indent

import (
	"regexp"
	"strings"
)

var lineExpression *regexp.Regexp
func init() {
	lineExpression = regexp.MustCompile("(?m)^(.+)$")
}

// Indent prepends the given indent to the given string
func Indent(input string, indent string, countOfTimes int) string {

	// Result pattern used to substitute into.
	indentation := strings.Repeat(indent, countOfTimes)
	resultPattern := []byte(indentation + "$1")

	indentedBytes := lineExpression.ReplaceAll([]byte(input), resultPattern)
	return string(indentedBytes)
}
