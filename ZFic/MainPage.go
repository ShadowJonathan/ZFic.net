package ZFic

import "net/http"

func MainPage() func(w http.ResponseWriter, r *http.Request) {
	return http.FileServer(http.Dir(webroot)).ServeHTTP
}
