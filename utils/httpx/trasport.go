package httpx

import (
	"context"
	"crypto/tls"
	"net/http"
)

func Context(ctx context.Context, insecureSkipVerify bool, baseTripper http.RoundTripper) http.RoundTripper {
	return &contextTransport{
		Context:            ctx,
		RoundTripper:       baseTripper,
		insecureSkipVerify: insecureSkipVerify,
	}
}

type contextTransport struct {
	context.Context
	http.RoundTripper
	insecureSkipVerify bool
}

func (tripper *contextTransport) Default(insecureSkipVerify bool) *http.Transport {
	transport := http.DefaultTransport.(*http.Transport)
	if insecureSkipVerify {
		if transport.TLSClientConfig == nil {
			transport.TLSClientConfig = &tls.Config{}
		}
		transport.TLSClientConfig.InsecureSkipVerify = true
	}
	return transport
}

func (tripper *contextTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if tripper.RoundTripper == nil {
		tripper.RoundTripper = tripper.Default(tripper.insecureSkipVerify)
	}
	if tripper.Context != nil {
		request = request.WithContext(tripper.Context)
	}
	return tripper.RoundTripper.RoundTrip(request)
}

var _ http.RoundTripper = (*contextTransport)(nil)
