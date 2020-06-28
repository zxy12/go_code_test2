package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	path := "/fsf/fss/xxx.pdf"

	fmt.Println((path))
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))
	newPath := fmt.Sprintf("%s/%s%s", filepath.Dir(path), filepath.Base(path), ".docx")
	fmt.Println(newPath)
	fmt.Println(filepath.Dir(path) + "/" + strings.TrimSuffix(filepath.Base(path), ".pdf") + ".docx")

	fmt.Println(strings.TrimPrefix(filepath.Ext(path), "."))

}
