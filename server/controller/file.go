package controller

import (
	"net/http"
	"strings"

	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/service"
)

type FileController struct {
	Logger      log.GoDropLogger
	FileService *service.FileService
	TrimSuffix  string
}

func (fc FileController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileId := strings.TrimPrefix(r.URL.Path, fc.TrimSuffix)
	switch r.Method {
	case http.MethodGet:
		err := fc.FileService.Download(r.Context(), fileId, w)
		if err != nil {
			fc.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		file, _, err := r.FormFile("file")
		if err != nil {
			fc.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = fc.FileService.Upload(r.Context(), fileId, file)
		if err != nil {
			fc.Logger.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
