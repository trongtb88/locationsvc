package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/trongtb88/locationsvc/src/business/entity"
	"github.com/trongtb88/locationsvc/src/common"
)

func Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !isValid(user, pass) {
			jsonErrResp := &entity.HTTPErrResp{
				Meta: entity.Meta{
					Path:       r.URL.String(),
					StatusCode: http.StatusUnauthorized,
					Status:     http.StatusText(http.StatusUnauthorized),
					Error: entity.ErrorMessage{
						Code:    http.StatusText(http.StatusUnauthorized),
						Message: http.StatusText(http.StatusUnauthorized),
					},
					Timestamp:  time.Now().Format(time.RFC3339),
				},
			}
			raw, _ := json.Marshal(&jsonErrResp)

			w.Header().Set(common.HttpHeaderContentType, common.HttpContentJSON)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(raw)
			return
		}
		handler(w, r)
	}
}

func isValid(username, password string) bool {
	configUsername := os.Getenv("Auth_Username")
	configPassword := os.Getenv("Auth_Password")
	return username == configUsername && password == configPassword
}
