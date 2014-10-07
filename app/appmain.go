package appmain

import (
	"net/http"

	"github.com/lancefisher/goappscaffold/handlers/guestbook"
	"github.com/lancefisher/goappscaffold/handlers/sign"
)

func init() {
	http.HandleFunc("/guest", guestbook.HandleFunc)
	http.HandleFunc("/sign", sign.HandleFunc)
}
