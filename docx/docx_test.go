package docx

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTxtToDocx(t *testing.T) {
	t.Log("done.")
	ss := `
	ff
	dd

	ee
	00
	1
	`
	str, err := StringToDocx(ss)
	if err != nil {
		t.Error(err)
	}

	d1 := []byte(str)
	err = ioutil.WriteFile("./test.docx", d1, 0644)
	fmt.Println(len(str))
	if err != nil {
		t.Error(err)
	}
}
