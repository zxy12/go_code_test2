package main

import "fmt"
import "strings"
import "flag"

var url = flag.String("u", "", "input url")

func main() {
	flag.Parse()
	str := *url
	bns, uri := parseBnsUrl(str)
	fmt.Println(bns, uri)

}

func parseBnsUrl(str string) (string, string) {
	if strings.HasPrefix(str, "bns://") {
		str = str[6:len(str)]
		if i := strings.Index(str, "///"); i > 0 {
			bns := str[0:i]
			uri := str[i+2 : len(str)]
			return bns, uri
		}
	}
	return "", ""
}
