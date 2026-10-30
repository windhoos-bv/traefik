// Package requestid is a Traefik middleware plugin that tags every request with a random id.
package requestid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

// Config is the plugin configuration.
type Config struct {
	HeaderName string `json:"headerName,omitempty"`
	// Keep an id that is already present on the incoming request (e.g. set by an upstream proxy).
	TrustIncoming bool `json:"trustIncoming,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{HeaderName: "X-Request-Id"}
}

// RequestID is the middleware.
type RequestID struct {
	next          http.Handler
	headerName    string
	trustIncoming bool
}

// New creates the middleware.
func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	headerName := config.HeaderName
	if headerName == "" {
		headerName = "X-Request-Id"
	}
	return &RequestID{next: next, headerName: headerName, trustIncoming: config.TrustIncoming}, nil
}

func (r *RequestID) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	id := req.Header.Get(r.headerName)
	if !r.trustIncoming || id == "" {
		id = newID()
	}
	// Modify the header map in place, so the access log (which holds a reference to it) sees the id.
	req.Header.Set(r.headerName, id)
	rw.Header().Set(r.headerName, id)
	r.next.ServeHTTP(rw, req)
}

// newID returns 32 hex characters, the same format as nginx's $request_id.
func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
