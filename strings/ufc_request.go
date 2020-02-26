package main

import "fmt"
import "strings"
import "flag"

var url = flag.String("u", "", "input url")

func main() {
	flag.Parse()
	str := *url
	UFC, uri := parseUFC(str)
	fmt.Println(UFC, uri)

}

func parseUFC(str string) (string, string) {
	if strings.HasPrefix(str, "ufc://") {
		str = str[6:len(str)]
		if i := strings.Index(str, "///"); i > 0 {
			UFC := str[0:i]
			uri := str[i+2 : len(str)]
			return UFC, uri
		}
	}
	return "", ""
}
