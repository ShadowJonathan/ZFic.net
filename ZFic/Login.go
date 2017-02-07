package ZFic

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(HandlePage(LP))
	} else if r.Method == http.MethodPost {
		if r.FormValue("status") == "bad" {
			w.Write(HandlePage(BLP))
		} else {
			HandleLogin(w, r)
		}
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	Sess, err := r.Cookie("zfsn")
	if err == nil {
		Kill(Sess.Value)
	}
	DeleteCookie(w, r, "zfun")
	DeleteCookie(w, r, "zfid")
	DeleteCookie(w, r, "zfsn")
	http.Redirect(w, r, "/", 307)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("UN")
	p := r.FormValue("PW")
	if u == "" || p == "" {
		http.Redirect(w, r, "/login/?status=bad", 307)
	} else {
		yus, U := GetUser(u, p)
		if !yus {
			http.Redirect(w, r, "/login/?status=bad", 307)
		} else {
			c := &http.Cookie{
				Name:  "zfun",
				Value: U.Username,
				Path:  "/",
			}
			http.SetCookie(w, c)
			GetNewSession(GSFI(U.ID), w)
			http.Redirect(w, r, "/u/"+GSFI(U.ID), 307)
			//Us, _ := json.Marshal(U)
			//w.Write(Us)
		}
	}
}
