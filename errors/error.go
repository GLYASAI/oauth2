package errors

import "errors"

// known errors
var (
	ErrInvalidRedirectURI   = errors.New("invalid redirect uri")
	ErrInvalidClientID      = errors.New("invalid client id")
	ErrInvalidClientSecret  = errors.New("invalid client secret")
	ErrInvalidAuthorizeCode = errors.New("invalid authorize code")
	ErrInvalidAccessToken   = errors.New("invalid access token")
	ErrInvalidRefreshToken  = errors.New("invalid refresh token")
	ErrExpiredAccessToken   = errors.New("expired access token")
	ErrExpiredRefreshToken  = errors.New("expired refresh token")
	ErrInvalidContentType   = errors.New("invalid content type")
)
