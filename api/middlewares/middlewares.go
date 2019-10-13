package middlewares

import(
	"errors"
	"github.com/nozgurozturk/startpage_server/api/auth"
	"github.com/nozgurozturk/startpage_server/api/utils"
	"net/http"
)

func JsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		err := auth.TokenValid(r)
		if err != nil{
			utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}