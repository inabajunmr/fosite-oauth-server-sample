package main

import (
	"log"
	"net/http"

	"github.com/inabajunmr/fosite-oauth-server-sample/oauth2"
)

func main() {
	http.HandleFunc("/oauth2/auth", oauth2.AuthorizationEndpoint)
	http.HandleFunc("/oauth2/token", oauth2.TokenEndpoint)
	http.HandleFunc("/oauth2/introspect", oauth2.IntrospectionEndpoint)
	log.Fatal(http.ListenAndServe(":3846", nil))
}
