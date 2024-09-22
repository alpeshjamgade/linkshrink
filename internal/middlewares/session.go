package middlewares

import (
	"context"
	"net/http"
	"shrinklink/internal/constants"
	"shrinklink/internal/utils"
)

type ISession interface {
	ValidateSession(next http.Handler) http.Handler
}

type Session struct {
	session ISession
}

func NewSession(session ISession) *Session { return &Session{session: session} }

func (s *Session) ValidateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := utils.GetUUID()
		ctx := context.WithValue(r.Context(), constants.TRACE_ID, traceId)

		// validate session
		_, err := r.Cookie("TOKEN")
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
