package ZFic

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	MP, err := ioutil.ReadFile(MPTP)
	if err != nil {
		panic(err)
	}
	MPMD, err := ioutil.ReadFile(fileroot + "/MP.md")
	MAINPAGE, err := template.New("MainPage").Parse(string(MP))
	MAINPAGE.Execute(w, template.HTML(string(GetMD(MPMD))))
	if err != nil {
		panic(err)
	}
	return
}
