package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var url = flag.String("u", "", "input download url")
var outfile = flag.String("o", "", "input output file")
var r = flag.Bool("r", false, "is follow redirect, default false")

func main() {
	flag.Parse()
	flag.Usage = func() {
		flag.PrintDefaults()
		os.Exit(2)
	}
	if *url == "" {
		flag.Usage()
		return
	}

	var httpreq *http.Request
	var httpres *http.Response
	var err error

	client := &http.Client{}
	if httpreq, err = http.NewRequest("GET", *url, nil); err != nil {
		panic("make request fail")
	}
	httpreq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	if !*r {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= 0 {
				return fmt.Errorf("cant redirect")
			}
			return nil
		}
	}

	httpres, err = client.Do(httpreq)
	if httpres == nil {
		panic("httpres nil")
	}
	fmt.Printf("%+v\n", httpres)
	fmt.Printf("response=%T,%+v,%+v,statusCode=%v\n", httpres.Header, httpres.Header, err, httpres.StatusCode)
	if err != nil {
		panic("request fail" + err.Error())
	}

	defer httpres.Body.Close()

	if *outfile == "" {
		return
	}
	filename := *outfile
	file, err := os.Create(filename)
	if err != nil {
		panic("create file fail" + err.Error())
	}
	defer file.Close()
	_, err = io.Copy(file, httpres.Body)
	if err != nil {
		panic("download fail")
	}

}
