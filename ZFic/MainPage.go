package ZFic

import (
	"io"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n") // yes, its just for testing purposes >.>
}
