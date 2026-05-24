package docs

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /openapi.json", OpenAPI)
	mux.HandleFunc("GET /docs", SwaggerUI)
}
