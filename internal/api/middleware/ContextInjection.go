package middleware

import (
	context2 "context"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/context"
	"github.com/yukitsune/chameleon/internal/api/modules"
	"github.com/yukitsune/chameleon/internal/api/responseWriterHelpers"
	"net/http"
)

type ContextInjection struct {
	container camogo.Container
}

func NewContainerInjectionMiddleware(container camogo.Container) *ContextInjection {
	return &ContextInjection{container: container}
}

func (m *ContextInjection) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Behaviour here is kinda weird:
		// - ctx needs the child container
		// - child container needs a context

		// So:
		// - The container gets the original context (without the container itself inside)
		// - The request gets the new context with the container

		// Meaning:
		// - No circular reference (ctr -> ctx -> ctr -> ctx -> ...)
		// - Container access from context limited to HTTP handlers

		cc, err := m.container.NewChildWith(func (cb camogo.ContainerBuilder) error {
			err := cb.RegisterFactory(func () context2.Context {
				return r.Context()
			}, camogo.SingletonLifetime)
			if err != nil {
				return nil
			}

			// Todo: Not sure i like registering modules this far down the line...
			// Todo: Experiment here, how would autofac do it?
			return cb.RegisterModule(&modules.MediatorHandlerModule{})
		})

		if err != nil {
			responseWriterHelpers.Error(w, err)
			return
		}

		ctx := context.WithContainer(r.Context(), cc)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
