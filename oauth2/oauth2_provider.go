package oauth2

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	"github.com/inabajunmr/fosite-oauth-server-sample/storage"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/token/jwt"
)

var secret = []byte("my super secret signing password")
var privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
var config = &compose.Config{}
var oauth2Provider = compose.Compose(
	config,
	storage.NewInMemoryStorage(defaultClient()),
	&compose.CommonStrategy{
		CoreStrategy:               compose.NewOAuth2HMACStrategy(config, secret, nil),
		OpenIDConnectTokenStrategy: compose.NewOpenIDConnectStrategy(config, privateKey),
		JWTStrategy: &jwt.RS256JWTStrategy{
			PrivateKey: privateKey,
		},
	},
	nil,

	compose.OAuth2ClientCredentialsGrantFactory,
	compose.OAuth2TokenIntrospectionFactory,
)

func defaultClient() fosite.Client {
	hasher := &fosite.BCrypt{WorkFactor: 6}
	secret, _ := hasher.Hash(context.TODO(), []byte("secret"))
	return &fosite.DefaultClient{
		ID:         "default-client",
		Secret:     secret,
		GrantTypes: []string{"client_credentials"},
	}
}
