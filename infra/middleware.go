package infra

import (
	"net/http"
)

type MiddlewareFactory struct {
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func InitMiddlewareFactory(middleware ...func(http.HandlerFunc) http.HandlerFunc) *MiddlewareFactory {
	var mws []func(http.HandlerFunc) http.HandlerFunc
	for _, m := range middleware {
		mws = append(mws, m)
	}
	return &MiddlewareFactory{
		Middlewares: mws,
	}
}

func (factory *MiddlewareFactory) Get(fun http.HandlerFunc) http.HandlerFunc {
	for _, mw := range factory.Middlewares {
		fun = mw(fun)
	}
	return fun
}
