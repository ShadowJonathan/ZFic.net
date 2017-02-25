package ZFic

import "net/http"

//noinspection GoUnusedExportedFunction
func Static() func(w http.ResponseWriter, r *http.Request) {
	return http.FileServer(http.Dir(ZFroot + "/static/")).ServeHTTP
}
