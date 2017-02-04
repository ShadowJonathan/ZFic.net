package main

import (
	"log"
	"net/http"

	"fmt"

	"../ZFic"
)

func main() {
	ZFIC, Sucess := ZFic.Load()
	if Sucess != nil {
		log.Fatal("Error Loading Server: " + Sucess.Error())
	}
	ZF := http.NewServeMux()
	ZF.HandleFunc("/", ZFic.MainPage)                                         // the front page
	ZF.HandleFunc("/a", ZFic.Archive)                                         // archive handler, aka search
	ZF.HandleFunc("/f/", ZFic.Fic)                                            // story handler, aka fanfiction
	ZF.HandleFunc("/static/", http.FileServer(http.Dir("../Zfic")).ServeHTTP) // the css and other shit >.> (probably images and stuff)
	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe(ZFIC.Address()+":"+ZFIC.Port(), ZF))
}
