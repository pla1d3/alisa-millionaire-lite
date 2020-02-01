package middlewares

import "net/http"

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		next.ServeHTTP(res, req)
	})
}
