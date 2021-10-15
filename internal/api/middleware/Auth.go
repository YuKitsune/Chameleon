package middleware

import (
	"github.com/yukitsune/chameleon/internal/api/context"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"net/http"
)

const (
	ApiKeyHeader = "X-Chameleon-Access-Token"
	AuthorizationHeader = "Authorization"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		container, err := context.Container(r.Context())
		if err != nil {
			responseWriterHelpers.Error(w, err)
			return
		}

		// Check for JWT
		authHeader := r.Header[AuthorizationHeader]
		if len(authHeader) > 1 {
			// Todo: Fail: Only 1 supported
		}
		if len(authHeader) == 1 {
			// Todo: Validate JWT
		}

		// Check for access token
		apiKeyHeader := r.Header[ApiKeyHeader]
		if len(apiKeyHeader) > 1 {
			responseWriterHelpers.BadRequestf(w, "only one %s is allowed", ApiKeyHeader)
			return
		}

		if len(apiKeyHeader) == 1 {
			apiKeyValue := apiKeyHeader[0]
			allow, err := container.ResolveWithResult(func (mediator mediator.Mediator) (bool, error) {
				res, err := mediator.Send(&model.CheckApiKeyRequest{Value: apiKeyValue})
				return res.(bool), err
			})

			if err != nil {
				responseWriterHelpers.Error(w, err)
				return
			}

			if allow.(bool) {
				next.ServeHTTP(w, r)
			}
		}

		responseWriterHelpers.Unauthorized(w)
	})
}
