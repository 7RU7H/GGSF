// You can edit this code!
// Click here and start typing.
package GGSF

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	//cmds:
	//hex -b Bad Characters
	//url encode
	//requote -b "\""
	//unicode crazy word to describe using weird characters
	//escape -p [PATTERN] -e [PATTERN TO ESPACE WITH ACCEPTS MUTLIPLE CHARACTERS]
	//lfi -e url -b "."
	//casechange - create all permuations for case

	//url -p
	//url -q
	//url -u path|query

	// Where QUOTE is equal to " ' `
	//requote -r "\QUOTE" -n "\QUOTE"
	fmt.Println(usage)

}

// goHelp/helpConvEnc.go
func convStrSliceNumToHex(s []string) byte {
	return hex.EncodeToString([]byte(str))
}

// Test to consider to achieve this - literal strings that have quotes inside
func requoteSanatizedInputStr(input, unwantedQuotes, desiredQuotes string) (error string) {
	return strings.Replace(input, unwantedQuotes, desiredQuotes, -1), error
}

// Test to consider to achieve this
func urlPathEncode(input string) (error string) {
	return url.PathEscape(input)
}

func urlPathEncode(input string) (error string) {
	return url.QueryEscape(input)
}

func urlPathEncode(input string) (error string) {
	return url.PathUnEscape(input)
}

func urlPathEncode(input string) (error string) {
	return url.QueryUnescape(input)
}

func charEscape(input, escapePattern, escapeSymbol string) (error string) {
	return strings.Replace(input, escapePattern, escapeSymbol, -1), error
}
