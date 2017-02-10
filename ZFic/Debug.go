package ZFic

import "net/http"
import "fmt"

// code in here should only be run when server is in debug mode

func Reinstallmembers(w http.ResponseWriter, r *http.Request) {
	LoadDataBase()
	w.Write([]byte("Reloaded database"))
	fmt.Println(r.RemoteAddr + " called reinstallmembers")
}

func Killallsessions(w http.ResponseWriter, r *http.Request) {
	count := DieDieDie()
	w.Write([]byte(GSFI(count) + " session(s) have been killed"))
	fmt.Println(r.RemoteAddr + " called killallmembers")
}
