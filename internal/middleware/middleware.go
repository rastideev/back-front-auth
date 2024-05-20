package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewareList ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {

		for i := len(middlewareList) - 1; i >= 0; i-- {
			middleware := middlewareList[i]
			next = middleware(next)
		}

		return next
	}
}
