package middleware

import (
	"fmt"
	"lets-go-book-2022/cmd/web/base"
	"log/slog"
	"net/http"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

type LogMiddleware struct {
	Logger *slog.Logger
}

func (m *LogMiddleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Info("Log Middleware", "RemoteAddr", r.RemoteAddr, "Proto", r.Proto, "Method", r.Method, "RequestURI", r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

type RecoverPanicMiddleware struct {
	App *base.Application
}

func (m *RecoverPanicMiddleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				m.App.ServerError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
