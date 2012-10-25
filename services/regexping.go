package services

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"strings"
)
//error: line 1 in file "/tmp/gff684122466" does not begin with "##gff-version" or "##gff-version""
// var rexp = regexp.MustCompile("^(warning)|(error): (.+) line ([:digit:]+) in file \"[^\"]*\"( (.*))$")
var rexp = regexp.MustCompile("(warning|error):[ ]?(.*) line (\\d+) in file (\"[^\"]*\")(?: (.*))?$")
func parseStderr (b *bytes.Buffer) ([]string, []string) {
	var errors []string
	var warnings []string
	for {
		line, err := b.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println("debug: error while regexping")
			return errors, warnings
		}
		line = strings.TrimSpace(line)

		match := rexp.FindStringSubmatch(line)
		if match != nil {
			highlighted := "<b>" + match[1] + "</b> " + match[2] + " <b>line " + match[3] + "</b> " + match[5]
			if match[1] == "error" {
				errors = append(errors, highlighted)
			} else if match[1] == "warning" {
				warnings = append(warnings, highlighted)
			}
		} else {
			if line != "" {
				errors = append(errors, line)
				log.Println("Unmatched line: "+ line)
			}
		}
		
		if err == io.EOF {
			log.Println("debug: end of input")
			break
		}
	}
	return errors, warnings
}

// func main () {
// 	r := ParseError([]byte("warning: skipping blank line 31 in file \"/tmp/gff262374079\""))
// 	for _,v := range r {
// 		log.Println(string(v))
// 	}
// 	r = ParseError([]byte("warning: seqid \"Ca21chr5_C_albicans_SC5314\" on line 3 in file \"/tmp/gff262374079\" has not been previously introduced with a \"##sequence-region\" line, create such a line automatically"))
// 	for _,v := range r {
// 		log.Println(string(v))
// 	}
// 	
// }