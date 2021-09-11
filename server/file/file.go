package file

import (
	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/service"
	"net/http"
	"strings"
)

type Handler struct {
	Logger      log.GoDropLogger
	FileService *service.FileService
	TrimSuffix  string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileId := strings.TrimPrefix(r.URL.Path, h.TrimSuffix)
	switch r.Method {
	case http.MethodGet:
		err := h.FileService.Download(r.Context(), fileId, w)
		if err != nil {
			h.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		file, _, err := r.FormFile("file")
		if err != nil {
			h.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = h.FileService.Upload(r.Context(), fileId, file)
		if err != nil {
			h.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
