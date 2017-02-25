package ZFic

import (
	"fmt"
	"net/http"
	"strings"
)

//noinspection GoUnusedExportedFunction
func AoZ(w http.ResponseWriter, r *http.Request) {

	page := strings.Replace(r.URL.Path, "/aoz/", "", 1)
	var aozroot = fileroot + "/AoZ/"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!!!!!!!!!!!!!!!!\nPANICKED:", r)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Critical error when displaying web-page: " + r.(error).Error()))
		}
	}()

	switch strings.ToLower(page) {
	case "login":
		w.Write(HandlePage(aozroot + "Login.icfg"))
		//nd
	case "badlogin":
		w.Write(HandlePage(aozroot + "badlogin.icfg"))
		//nd
	case "signup":
		w.Write(HandlePage(aozroot + "Signup.icfg"))
		//nd
	case "user":
		w.Write(HandlePage(aozroot + "User.icfg"))
		//nd
	case "settings":
		w.Write(HandlePage(aozroot + "Settings.icfg"))
		//nd
	case "story":
		w.Write(HandlePage(aozroot + "Story.icfg"))
		//nd
	case "reviewpage":
		w.Write(HandlePage(aozroot + "Review.icfg"))
		//nd
	case "":
		w.Write(HandlePage(aozroot + "Mainpage.icfg"))
		//nd
	default:
		w.Write(HandlePage(aozroot + "Mainpage.icfg"))
	}
}
