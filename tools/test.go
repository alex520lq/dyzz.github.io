package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	files, _ := ioutil.ReadDir("../images")
	baseUrl := "https://media.githubusercontent.com/media/dyzz/dyzz.github.io/master/images/"
	indexUrl := "![%s](pages/tmp/%s.md)\n"
	index := ""
	imgTempl := "![](%s%s.png)"
	for _, f := range files {
		fname := f.Name()
		extension := filepath.Ext(fname)
		name := fname[0 : len(fname)-len(extension)]
		data := fmt.Sprintf(imgTempl, baseUrl, name)
		ioutil.WriteFile("../zh_CN/pages/tmp/"+name+".md", []byte(data), 0644)
		index += fmt.Sprintf(indexUrl, name, name)
	}
	ioutil.WriteFile("../zh_CN/pages/tmp/index.md", []byte(index), 0644)
}
