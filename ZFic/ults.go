package ZFic

import (
	"bytes"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

func ProduceMainPage() {
	MP, err := ioutil.ReadFile(MPTP)
	if err != nil {
		panic(err)
	}
	MPMD, err := ioutil.ReadFile(fileroot + "/MP.md")
	if err != nil {
		return
	}
	PAGE := bytes.Replace(MP, []byte("MAINPAGEDATA"), GetMD(MPMD), 1)
	err = ioutil.WriteFile(webroot+"/index.html", PAGE, 9001)
	if err != nil {
		panic(err)
	}
}

func GetMD(input []byte) []byte {
	return blackfriday.MarkdownCommon(input)
}
