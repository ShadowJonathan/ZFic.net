package main

import (
	"log"
	"net/http"

	"fmt"

	"../ZFic"
)

func main() {
	Sucess, y := ZFic.Load()
	if !Sucess {
		fmt.Println("Error Loading Server: " + y)
	}
	ZF := http.NewServeMux()
	ZF.HandleFunc("/", ZFic.MainPage)
	ZF.HandleFunc("/a", ZFic.Archive)
	ZF.HandleFunc("/f", )
	log.Fatal(http.ListenAndServe(":80", ZF))
}
