package storage

import (
	"context"
	"sync"
	"time"

	"github.com/ory/fosite"
)

type InMemoryStorage struct {
	clients       sync.Map
	accessTokens  sync.Map
	refreshTokens sync.Map
}

func NewInMemoryStorage(defaultClient fosite.Client) *InMemoryStorage {
	s := InMemoryStorage{
		clients:       sync.Map{},
		accessTokens:  sync.Map{},
		refreshTokens: sync.Map{},
	}
	s.CreateClient(context.TODO(), defaultClient)
	return &s
}

func (s *InMemoryStorage) CreateClient(_ context.Context, client fosite.Client) {
	s.clients.Store(client.GetID(), client)
}

func (s *InMemoryStorage) GetClient(_ context.Context, id string) (fosite.Client, error) {
	client, ok := s.clients.Load(id)
	if ok {
		return client.(fosite.Client), nil
	}

	return nil, fosite.ErrNotFound
}

func (s *InMemoryStorage) ClientAssertionJWTValid(_ context.Context, jti string) error {
	return nil
}

func (s *InMemoryStorage) SetClientAssertionJWT(_ context.Context, jti string, exp time.Time) error {
	return nil
}

func (s *InMemoryStorage) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	s.accessTokens.Store(signature, request)
	return nil
}

func (s *InMemoryStorage) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	at, ok := s.accessTokens.Load(signature)
	if ok {
		return at.(fosite.Requester), nil
	}

	return nil, fosite.ErrNotFound
}

func (s *InMemoryStorage) DeleteAccessTokenSession(ctx context.Context, signature string) (err error) {
	s.accessTokens.Delete(signature)
	return nil
}

func (s *InMemoryStorage) CreateAuthorizeCodeSession(ctx context.Context, code string, request fosite.Requester) (err error) {
	return nil
}

func (s *InMemoryStorage) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (request fosite.Requester, err error) {
	return nil, nil
}

func (s *InMemoryStorage) InvalidateAuthorizeCodeSession(ctx context.Context, code string) (err error) {
	return nil
}

func (s *InMemoryStorage) CreateRefreshTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	s.refreshTokens.Store(signature, request)
	return nil
}

func (s *InMemoryStorage) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	at, ok := s.refreshTokens.Load(signature)
	if ok {
		return at.(fosite.Requester), nil
	}

	return nil, fosite.ErrNotFound
}

func (s *InMemoryStorage) DeleteRefreshTokenSession(ctx context.Context, signature string) (err error) {
	s.refreshTokens.Delete(signature)
	return nil
}
