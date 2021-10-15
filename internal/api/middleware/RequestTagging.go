package middleware

import (
	"github.com/google/uuid"
	"github.com/yukitsune/chameleon/internal/api/context"
	"net/http"
)

const RequestIdHeaderKey = "X-Request-Id"

func RequestTagging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.New().String()

		w.Header().Set(RequestIdHeaderKey, reqId)
		ctx := context.WithRequestId(r.Context(), reqId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}