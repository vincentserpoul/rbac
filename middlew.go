package rbac

import (
	"log"
	"net/http"
)

// Auth is an interface to be implemented by any rbac source, ie jsonfile rbac
type Auth interface {
	IsUserAuthorized(userID string, action string) (bool, error)
}

// InterposeRBAC returns a Handler that checks if user is logged in. Writes a http.StatusUnauthorized
func InterposeRBAC(authoriz Auth, funcGetUserID func(*http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			userID := funcGetUserID(req)
			action := req.Method + "_" + req.RequestURI
			isAuthorized, err := authoriz.IsUserAuthorized(userID, action)

			if !isAuthorized {
				log.Println(err)
				http.Error(res, "Not Authorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(res, req)
		})
	}
}
