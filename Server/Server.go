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
	ZF.HandleFunc("/", ZFic.MainPage()) // the front page
	ZF.HandleFunc("/a", ZFic.Archive)   // archive handler, aka search
	ZF.HandleFunc("/f", ZFic.Fic)       // story handler, aka fanfiction

	ZF.HandleFunc("/static", ZFic.Static()) // the css and other shit >.> (probably images and stuff)
	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8080", ZF))
}
