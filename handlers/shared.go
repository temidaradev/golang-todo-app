package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(c fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")

	return component.Render(c.Context(), c.Response().BodyWriter())
}
