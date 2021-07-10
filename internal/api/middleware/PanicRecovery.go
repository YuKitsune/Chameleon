package middleware

import (
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"github.com/yukitsune/chameleon/internal/log"
	"net/http"
	"runtime"
)

type PanicRecovery struct {
	container camogo.Container
}

func NewPanicRecovery(container camogo.Container) *PanicRecovery {
	return &PanicRecovery{container: container}
}

func (m *PanicRecovery) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				_ = m.container.Resolve(func(logger log.ChameleonLogger) {
					if logger.GetLevel() == log.TraceLevel {
						logger.Errorf("Recovering from panic: %v\n%s\n", err, buf)
					} else {
						logger.Errorf("Recovering from panic: %v\n", err)
					}
				})

				responseWriterHelpers.WriteEmptyResponse(w, http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
