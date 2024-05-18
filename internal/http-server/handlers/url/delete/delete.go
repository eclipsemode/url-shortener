package delete

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	responseApi "url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"
)

type Response struct {
	responseApi.Response
	Alias string `json:"alias,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=URLDelete
type URLDelete interface {
	DeleteUrl(alias string) error
}

func New(log *slog.Logger, delete URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, responseApi.Error("invalid request"))

			return
		}

		err := delete.DeleteUrl(alias)
		if errors.Is(err, storage.ErrNotFound) {
			log.Info("resource not found", "alias", alias)

			render.JSON(w, r, responseApi.Error("resource not found"))

			return
		}

		if err != nil {
			log.Error("failed to delete", "alias", alias, "err", err)

			render.JSON(w, r, responseApi.Error("failed to delete"))

			return
		}

		log.Info("successfully deleted url", "alias", alias)

		responseOk(w, r, alias)
	}

}

func responseOk(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: responseApi.Ok(),
		Alias:    alias,
	})
}
