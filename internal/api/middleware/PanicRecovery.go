package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/internal/api/context"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"net/http"
	"runtime"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				// Todo: It'd be nice to not have to fetch the container here
				container, _ := context.Container(r.Context())
				if container != nil {
					_ = container.Resolve(func(logger *logrus.Logger) {
						logger.Errorf("Recovering from panic: %v\n%s\n", err, buf)
					})
				}

				responseWriterHelpers.EmptyError(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
