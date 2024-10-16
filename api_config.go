package main

import (
	"net/http"
)

type apiConfig struct {
	fileserverHits int
	jwtSecret      string
	polkaApiKey    string
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++

		next.ServeHTTP(w, r)
	})
}
