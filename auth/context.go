package auth

import (
	"context"
)

type credentialsContextKey struct{}

type tokenTypeKey struct{}

// WithCredentials 认证Ctx
func WithCredentials(ctx context.Context, cred *Credentials) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, credentialsContextKey{}, cred)
}

// WithCredentialsType 带Type认证Ctx
func WithCredentialsType(ctx context.Context, cred *Credentials, t TokenType) context.Context {
	ctx = WithCredentials(ctx, cred)
	return context.WithValue(ctx, tokenTypeKey{}, t)
}
