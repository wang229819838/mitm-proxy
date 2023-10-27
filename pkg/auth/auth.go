// pkg/auth.go

package auth

import (
	"mitm-proxy/pkg/db"
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc, dbPath, realm string) http.HandlerFunc {
	database, err := db.OpenDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !db.ValidateUser(database, user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}
		handler(w, r)
	}
}
