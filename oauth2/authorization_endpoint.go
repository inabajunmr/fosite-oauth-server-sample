package oauth2

import (
	"net/http"

	"github.com/ory/fosite"
)

func AuthorizationEndpoint(rw http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	ar, err := oauth2Provider.NewAuthorizeRequest(ctx, req)
	if err != nil {
		oauth2Provider.WriteAuthorizeError(rw, ar, err)
		return
	}

	if req.URL.Query().Get("username") == "" {
		rw.Write([]byte(`set username as query parameter.`))
		return
	}

	mySessionData := &fosite.DefaultSession{
		Username: req.Form.Get("username"),
	}

	response, err := oauth2Provider.NewAuthorizeResponse(ctx, ar, mySessionData)
	if err != nil {
		oauth2Provider.WriteAuthorizeError(rw, ar, err)
		return
	}

	oauth2Provider.WriteAuthorizeResponse(rw, ar, response)
}
