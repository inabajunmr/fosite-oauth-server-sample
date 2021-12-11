package oauth2

import (
	"log"
	"net/http"

	"github.com/ory/fosite"
)

func TokenEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	accessRequest, err := oauth2Provider.NewAccessRequest(ctx, req, new(fosite.DefaultSession))
	if err != nil {
		log.Println(err.Error())
		oauth2Provider.WriteAccessError(rw, accessRequest, err)
		return
	}

	response, err := oauth2Provider.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		log.Println(err.Error())
		oauth2Provider.WriteAccessError(rw, accessRequest, err)
		return
	}

	oauth2Provider.WriteAccessResponse(rw, accessRequest, response)
}
