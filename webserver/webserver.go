package webserver

import (
	"fmt"
	"httpredirecter/redirection"
	"net/http"
)

func CreateWebserver(port uint16, redirections []redirection.Redirection) error {
	for _, r := range redirections {
		http.Handle(r.From, http.RedirectHandler(r.To, http.StatusTemporaryRedirect))
	}

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
