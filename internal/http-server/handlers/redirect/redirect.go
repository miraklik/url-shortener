package redirect

import (
	"errors"
	"log/slog"
	resp "main/internal/lib/api/response"
	"main/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log := log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("invalid request"))

			return
		}

		resURL, err := urlGetter.GetUrl(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Error("url not found", "alias", alias)

			render.JSON(w, r, resp.Error("not founded"))

			return
		}
		if err != nil {
			log.Error("failed to get url", "alias", alias, "err", err)

			render.JSON(w, r, resp.Error("failed to get url"))

			return
		}

		log.Info("success", "alias", alias, "url", resURL)

		http.Redirect(w, r, resURL, http.StatusFound)
	}
}
