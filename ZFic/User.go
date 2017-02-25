package ZFic

import "net/http"

import "strings"
import "bytes"

//noinspection GoUnusedExportedFunction
func UserPage(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	if page == "/u/" {
		http.Redirect(w, r, "/", 307)
	} else {
		page = strings.Replace(page, "/u/", "", 1)
		GetUserPage(page, w, r)
	}
}

func GetUserPage(user string, w http.ResponseWriter, r *http.Request) {
	page := HandlePage(UP)
	yus, U := GetUserSimple(user)
	if yus {
		page = bytes.Replace(page, []byte("USER"), []byte(U.Name), 1)
		page = bytes.Replace(page, []byte("BIO"), []byte(U.Bio.Content), 1)
		page = bytes.Replace(page, []byte("BTIME"), []byte(GetDateTime(U.Bio.Lastupdated)+":"), 1)
		if U.Avatar != "" {
			page = bytes.Replace(page, []byte("UAVATAR"), []byte("/pic/"+GSFI(U.ID)+"/"+U.Avatar+".jpg"), 1)
		} else {
			//default avatar shit
		}
		w.Write(page)
	} else {
		HandleError(404, w)
	}
}

//noinspection GoUnusedExportedFunction
func UserSettingsPage(w http.ResponseWriter, r *http.Request) {
	w.Write(HandlePage(SP))
	/*page := HandlePage(SP)
	U, yus := GetFromSession(r)
	if yus {
		page = bytes.Replace(page, []byte("NAMENAME"), []byte(U.Name), 1)
		if U.Nick != "" {
			page = bytes.Replace(page, []byte("NICKNAME"), []byte(U.Nick), 1)
		} else {
			page = bytes.Replace(page, []byte("NICKNAME"), []byte("No Nickname set"), 1)
		}
		// page = bytes.Replace(page, []byte("BIOCHANGE"), []byte(), 1)
		if U.Avatar != "" {
			page = bytes.Replace(page, []byte("UAVATAR"), []byte("/pic/"+GSFI(U.ID)+"/"+U.Avatar+".jpg"), 1)
		} else {
			//default avatar shit
		}
		w.Write(page)
	} else {
		HandleError(404, w)
	}*/
}
