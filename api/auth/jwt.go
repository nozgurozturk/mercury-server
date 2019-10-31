package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/nozgurozturk/mercury-server/api/models"
	"github.com/nozgurozturk/mercury-server/api/utils"
	"net/http"
	"os"
	"strings"
)

func JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		withOutAuth := []string{"/signup", "/login", "/"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range withOutAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message(false, "Missing auth token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, http.StatusForbidden, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, http.StatusForbidden, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, http.StatusForbidden, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, http.StatusForbidden, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	});
}
