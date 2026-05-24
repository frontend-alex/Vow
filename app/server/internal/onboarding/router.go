package onboarding

import "net/http"

func Routes(mux *http.ServeMux, handler Handler, protection func(http.Handler) http.Handler) {
	mux.Handle("POST /v1/api/onboarding/start", protection(http.HandlerFunc(handler.Start)))
	mux.Handle("POST /v1/api/onboarding/complete", protection(http.HandlerFunc(handler.Complete)))
}
