package onboarding

import "net/http"

func Routes(mux *http.ServeMux, handler Handler) {
	mux.HandleFunc("POST /v1/api/onboarding/start", handler.Start)
	mux.HandleFunc("POST /v1/api/onboarding/complete", handler.Complete)
}
