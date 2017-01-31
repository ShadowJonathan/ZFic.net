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
	ZF.HandleFunc("/", ZFic.MainPage())
	ZF.HandleFunc("/a", ZFic.Archive)
	ZF.HandleFunc("/f", ZFic.Fic)

	ZF.HandleFunc("/static", ZFic.Static())

	log.Fatal(http.ListenAndServe(":8080", ZF))
}
