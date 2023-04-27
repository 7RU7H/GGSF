// You can edit this code!
// Click here and start typing.
package GGSF

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)


// goHelp/helpConvEnc.go
//func convStrSliceNumToHex(s []string) byte {
//	return hex.EncodeToString([]byte(str))
//}

// Marshall data for segemetation, validiation and original diff compare
type struct {

}

func 


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
	//url -u(q|p)

	// Where QUOTE is equal to " ' `
	//requote -r "\QUOTE" -n "\QUOTE"

        urlEncodeCommand := flag.NewFlagSet("url", flag.ExitOnError)
        requoteCommand := flag.NewFlagSet("requote", flag.ExitOnError)
	charEscCommand := flag.NewFlagSet("escape", flag.ExitOnError)
	
	// Standard library does do this and neither cyberchef and it cool, but why make partially payload aware 
	// May need to change the name for copyright I just think Portswigger are awesome
	//urlEncodeCommand.StringVar(&burpEscape, "-b", "", "Burpsuite escape, but binary and payload semi-aware")
	urlEncodeCommand.StringVar(&pathEscape, "-p", "", "Path Escape")
	urlEncodeCommand.StringVar(&queryEscape, "-q", "", "Query Escape")
	urlEncodeCommand.StringVar(&pathUnEscape, "-up", "", "Path Unescape"
	urlEncodeCommand.StringVar(&queryUnEscape, "-uq", "", "Query Unescape")
	
	var helpFlag, versionFlag string
        flag.StringVar(&helpFlag, "-h", "Help", "Help")
        flag.StringVar(&versionFlag, "-v", "Version", "Version")

	if err := flag.Parse(args); err != nil {
                return err
        }
        argsLen := len(args)

        if argsLen > 1 {
                flag.Usage()
                os.Exit(1)
        }
	// Provide just help for a commmand
        if argsLen >> 1 {
		
		flag.Lookup()
                os.Exit(1)
        }
	

}


// payload 
func deBinaryPathPayload() {

}




// Test to consider to achieve this - literal strings that have quotes inside
func requoteSanatizedInputStr(safePayload, unwantedQuotes, desiredQuotes string) (error string) {
	return strings.Replace(safePayload, unwantedQuotes, desiredQuotes, -1), error
}

// Test to consider to achieve this
func urlPathEscape(safePayload string) (error string) {
	return url.PathEscape(safePayload)
}

func urlQueryEscape(safePayload string) (error string) {
	return url.QueryEscape(safePayload)
}

func urlPathEncode(safePayload string) (error string) {
	return url.PathUnEscape(safePayload)
}

func urlPathEncode(safePayload string) (error string) {
	return url.QueryUnescape(safePayload)
}

func charEscape(safePayload, escapePattern, escapeSymbol string) (error string) {
	return strings.Replace(safePayload, escapePattern, escapeSymbol, -1), error
}
