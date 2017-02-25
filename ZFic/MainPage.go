package ZFic

import (
	"net/http"
)

//noinspection GoUnusedExportedFunction
func MainPage(w http.ResponseWriter, r *http.Request) {
	data := HandlePage(MP)
	w.Write(data)
	//http.ServeFile(w, r, Mainpagehtml)
}
