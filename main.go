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



// Store Metadata for testing, empirical comparisons and not cluttering main  
type metadata struct {
	badchars string
	cmd string
	outputOption int // 0 stdout; 1 file
	outputFilePath string
}

// Marshall data for segemetation, validiation and original diff compare
type data struct {
// stdin record for comparativity test

// Safe String Character Map 
	safeStringCharMap map

// Byte version of stdin
	bytePayload byte
// For section by section micro transformation
	ptrTransformStart int
	prtTransFormEnd int
}

// For actual result and intended result from metadata logic
type result struct {

}




// On initialisation covert to golang maliable bytes
func (data *d) convertStdinToStore() error {
	

}

func (data *d) initSafeStringCharMap() error {

	d.safeStringCharMap := make(map[int]int)
	// Update goHelp with old golang map best practices
	// !! 
	for i,char := range d.safeStringCharMap { 
		d.safeStringCharMap[i] = strconv.Atoi(char)
	}
		
}






// Result 
func (result *r) () {

}





func main() {
	//cmds:
	//hex -b Bad Characters
	//url encode
	//requote -b "\""
	//unicode crazy word to describe using weird characters
	//escape -p [PATTERN] -e [PATTERN TO ESPACE WITH ACCEPTS MUTLIPLE CHARACTERS]
	//lfi -e url -b "."
	//recase - create all permuations for case

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
	urlEncodeCommand.StringVar(&payload, "-x", "", "Declare payload, must be final flag")
	
	// Requote
	requoteCommand.StringVar(&payload, "-x", "", "Declare payload, must be final flag")
	
	// Escape with 
	charEscCommand.StringVar(&payload, "-x", "", "Declare payload, must be final flag")

	
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

	// Positional arguements 
	// After all flags is the payload


	// Perform all checks and store for later checks
	currentMetadata := metadata{
	}
	
	// If Metadata checks do not throw incompatiblity errors convert data  
	currentData := data{

	}


	// Payload Transformation starts


	// Result construction and printing to 
}


// payload 
func deBinaryPathPayload() {

}

// STDIN handling
// Alphanumeric or Special chars
// ' ` "  


// Non-standard Characters (later)


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
