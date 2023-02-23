package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	audience string
	domain   string
)

func setAuth0Variables() {
	audience = os.Getenv("AUTH0_API_IDENTIFIER")
	domain = os.Getenv("AUTH0_DOMAIN")
}

// MiddleWare function to handle authentication
func authRequired(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(r)

		if err != nil {
			//json.NewEncoder(w).Encode("Can't Authenticate")
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
