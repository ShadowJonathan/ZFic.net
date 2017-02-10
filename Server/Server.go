package main

import (
	"log"
	"net/http"

	"fmt"

	"io/ioutil"

	"strconv"

	"../ZFic"
)

func main() {
	ZFIC, Sucess := ZFic.Load()
	if Sucess != nil {
		log.Fatal("Error Loading Server: " + Sucess.Error())
	}
	ZF := http.NewServeMux()
	ZF.HandleFunc("/", ZFic.MainPage)                                         // the front page
	ZF.HandleFunc("/a", ZFic.Archive)                                         // archive and search
	ZF.HandleFunc("/u/", ZFic.UserPage)                                       // user pages
	ZF.HandleFunc("/login/", ZFic.Login)                                      // login page
	ZF.HandleFunc("/logout/", ZFic.LogOut)                                    // logout page (duh)
	ZF.HandleFunc("/signup/", ZFic.SignUp)                                    // signup page
	ZF.HandleFunc("/settings/", ZFic.UserSettingsPage)                        // settings page
	ZF.HandleFunc("/f/", ZFic.Fic)                                            // story handler, aka fanfiction
	ZF.HandleFunc("/static/", http.FileServer(http.Dir("../Zfic")).ServeHTTP) // css and fonts
	ZF.HandleFunc("/pic/", http.FileServer(http.Dir("../Zfic")).ServeHTTP)    // images
	fmt.Println("Starting server...")

	if Debug() {
		fmt.Println("Started debug...")
		ZF.HandleFunc("/reinstallmembers", ZFic.Reinstallmembers)
		ZF.HandleFunc("/killallsessions", ZFic.Killallsessions)
	}

	if Show() {
		ZF.HandleFunc("/show/", ZFic.ShowRoom)
	}

	log.Fatal(http.ListenAndServe(ZFIC.Address()+":"+ZFIC.Port(), ZF))
}

func Debug() bool {
	file, err := ioutil.ReadFile("debug")
	if err != nil {
		return false
	} else {
		b, err := strconv.ParseBool(string(file))
		if err != nil {
			return false
		}
		return b
	}
}

func Show() bool {
	file, err := ioutil.ReadFile("show")
	if err != nil {
		return false
	} else {
		b, err := strconv.ParseBool(string(file))
		if err != nil {
			return false
		}
		return b
	}
}
