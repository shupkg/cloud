package sts

import (
	"context"
)

type Service interface {
	GetFederationToken(ctx context.Context, options GetFederationTokenOptions) (result GetFederationTokenResult, err error)
}

