package middleware

import (
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/context"
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

		// Todo: Camogo needs to take a new dependency (r.Context()) here
		cc := m.container.NewChild()
		ctx := context.WithContainer(r.Context(), cc)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
