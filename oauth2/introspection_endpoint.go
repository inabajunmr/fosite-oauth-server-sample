package oauth2

import (
	"log"
	"net/http"

	"github.com/ory/fosite"
)

func IntrospectionEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ir, err := oauth2Provider.NewIntrospectionRequest(ctx, req, new(fosite.DefaultSession))
	if err != nil {
		log.Printf("Error occurred in NewIntrospectionRequest: %+v", err)
		oauth2Provider.WriteIntrospectionError(rw, err)
		return
	}
	oauth2Provider.WriteIntrospectionResponse(rw, ir)
}
