package middlewares

import (
	"net/http"
	"pink/app/errors"
	modelrepos "pink/domains/model_repos"
)

// AuthMiddleware struct
type AuthMiddleware struct {
}

// AuthAPIToken funnc
func (auth *AuthMiddleware) AuthAPIToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiToken := r.Header.Get("X-Api-Token")
		userRepo := modelrepos.NewUserRepo()
		if userRepo.Authenticate(apiToken) {
			next.ServeHTTP(w, r)
		} else {
			handlerError := errors.HandlersError{
				Writer:  w,
				Message: "Auth",
				Status:  http.StatusForbidden,
			}
			handlerError.AuthError()
		}
	})
}
