package ZFic

import (
	"fmt"
	"net/http"
	"strings"
)

//noinspection GoUnusedExportedFunction
func ShowRoom(w http.ResponseWriter, r *http.Request) {
	page := strings.Replace(r.URL.Path, "/show/", "", 1)
	var showroot = fileroot + "/show/"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!!!!!!!!!!!!!!!!\nPANICKED:", r)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Critical error when displaying web-page: " + r.(error).Error()))
		}
	}()

	switch strings.ToLower(page) {
	case "login":
		w.Write(HandlePage(showroot + "Login.icfg"))
		//nd
	case "badlogin":
		w.Write(HandlePage(showroot + "badlogin.icfg"))
		//nd
	case "signup":
		w.Write(HandlePage(showroot + "Signup.icfg"))
		//nd
	case "user":
		w.Write(HandlePage(showroot + "User.icfg"))
		//nd
	case "settings":
		w.Write(HandlePage(showroot + "Settings.icfg"))
		//nd
	case "story":
		w.Write(HandlePage(showroot + "Story.icfg"))
		//nd
	case "mainpage":
		w.Write(HandlePage(showroot + "Mainpage.icfg"))
		//nd
	case "reviewpage":
		w.Write(HandlePage(showroot + "Review.icfg"))
		//nd
	case "":
		w.Write(HandlePage(showroot + "Initial.icfg"))
		//nd
	}
	//forum, maybe?
}
