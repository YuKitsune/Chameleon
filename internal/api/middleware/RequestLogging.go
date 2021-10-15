package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/context"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"net/http"
	"time"
)

// our http.ResponseWriter implementation
type loggingResponseWriterDecorator struct {
	http.ResponseWriter
	statusCode int
}

func (r *loggingResponseWriterDecorator) WriteHeader(statusCode int) {

	// write status code using original http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)

	// capture status code
	r.statusCode = statusCode
}

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqId, err := context.RequestId(r.Context())
		if err != nil {
			responseWriterHelpers.Error(w, err)
			return
		}

		ctr, err := context.Container(r.Context())
		if err != nil {
			responseWriterHelpers.Error(w, err)
			return
		}

		// wrap original http.ResponseWriter
		rwd := loggingResponseWriterDecorator{
			ResponseWriter: w,
			statusCode: 0,
		}

		start := time.Now()
		next.ServeHTTP(&rwd, r)
		duration := time.Since(start)

		_ = ctr.Resolve(func (log *logrus.Logger) {
			reqLogger := log.WithField("requestId", reqId).
				WithField("method", r.Method).
				WithField("path", r.URL.Path).
				WithField("status", rwd.statusCode).
				WithField("duration", duration)

			reqLogger.Infoln("")
		})
	})
}