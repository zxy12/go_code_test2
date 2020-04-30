package main

import "fmt"
import "encoding/json"
import "net/url"

func main() {
	str := `%7B%22custom_content%22%3A%7B%22title%22%3A%22%E8%BD%AC%E6%8D%A2%E6%88%90%E5%8A%9F%22%2C%22desc%22%3A%22%E5%B7%B2%E6%88%90%E5%8A%9F%E8%BD%AC%E6%8D%A2%E4%B8%BAword%E6%96%87%E6%A1%A3%E2%80%9Cyyy.pdf%E2%80%9D%E3%80%82%22%2C%22thumbnail%22%3A%22https%3A%2F%2Fstaticsns.cdn.bcebos.com%2Famis%2F2020-2%2F1582877968959%2Fic_word.png%22%2C%22url%22%3A%22%22%2C%22trans_type%22%3A1%2C%22link_type%22%3A4%2C%22pdf_path%22%3A%22%2FRTC+2019+PDF+-+%E5%85%B1%E4%BA%AB%E7%89%88%2F10+%E6%9C%88+25+%E6%97%A5%2F%E5%A4%A7%E5%89%8D%E7%AB%AF%E5%BA%94%E7%94%A8%E5%BC%80%E5%8F%91%E6%8A%80%E6%9C%AF%E4%B8%93%E5%9C%BA%2F%E6%9D%A8%E5%B0%9A%E6%9E%97+-+%E7%BE%8E%E5%9B%A2+-+RTC+2019%E5%AE%9E%E6%97%B6%E4%BA%92%E8%81%94%E7%BD%91%E5%A4%A7%E4%BC%9A%E6%BC%94%E8%AE%B2.pdf%22%2C%22doc_meta%22%3A%7B%22path%22%3A%22%2FProcess.docx%22%2C%22md5%22%3A%22da23fec25o9a0593ae05aaa8ee39e657%22%2C%22size%22%3A42179%2C%22fs_id%22%3A249394312673375%2C%22category%22%3A4%7D%2C%22trans_status%22%3A0%2C%22clienttype%22%3A0%2C%22is_hidden%22%3A0%7D%7D`
	str, _ = url.QueryUnescape(str)
	fmt.Println(str)
	var data interface{}
	err := json.Unmarshal([]byte(str), &data)
	fmt.Printf("[%v], %+v", err, data)
	fmt.Println("vim-go")
}
