package routes

import (
	"encoding/json"
	"example/errors"
	"example/internal/core/application"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
)

type requestHandler[T any] func(application.Application, http.ResponseWriter, *http.Request) (*T, errors.AppError)

func HandleRequest[T any](application application.Application, handler requestHandler[T]) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var response *T
		var handlerError error

		defer func() {
			if exception := recover(); exception != nil {
				stk := debug.Stack()
				application.Logger().Error("Stack trace for exception: %v \n %v", exception, string(stk))
				if _, file, line, ok := runtime.Caller(1); ok {
					application.Logger().Error("Recovered from panic in file %s at line %d: %v\n", file, line, exception)
				} else {
					application.Logger().Error("Recovered from panic but couldn't retrieve file name and line number")
				}
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if handlerError != nil {
				if appError, ok := handlerError.(errors.AppError); ok {
					writer.WriteHeader(int(appError.Status()))
					// json.NewEncoder(writer).Encode(appError.Json())
				} else {
					application.Logger().Error(fmt.Sprintf("API Request error: %v", handlerError))
					http.Error(writer, handlerError.Error(), http.StatusInternalServerError)
				}
				return
			}
			json.NewEncoder(writer).Encode(map[string]any{
				"data": response,
			})
		}()
		response, handlerError = handler(application, writer, request)
	}
}

func Post(router *http.ServeMux, url string, handler http.HandlerFunc) {
	url = http.MethodPost + " " + url
	router.HandleFunc(url, handler)
}

func AuthPost(router *http.ServeMux, url string, handler http.HandlerFunc, secret string) {
	authHandler := JWTMiddleware(secret, handler)
	url = http.MethodPost + " " + url
	router.HandleFunc(url, authHandler)
}

func Get(router *http.ServeMux, url string, handler http.HandlerFunc) {
	url = http.MethodGet + " " + url
	router.HandleFunc(url, handler)
}

func AuthGet(router *http.ServeMux, url string, handler http.HandlerFunc, secret string) {
	authHandler := JWTMiddleware(secret, handler)
	url = http.MethodGet + " " + url
	router.HandleFunc(url, authHandler)
}

func Put(router *http.ServeMux, url string, handler http.HandlerFunc) {
	url = http.MethodPut + " " + url
	router.HandleFunc(url, handler)
}

func AuthPut(router *http.ServeMux, url string, handler http.HandlerFunc, secret string) {
	authHandler := JWTMiddleware(secret, handler)
	url = http.MethodPut + " " + url
	router.HandleFunc(url, authHandler)
}

func Patch(router *http.ServeMux, url string, handler http.HandlerFunc) {
	url = http.MethodPatch + " " + url
	router.HandleFunc(url, handler)
}

func AuthPatch(router *http.ServeMux, url string, handler http.HandlerFunc, secret string) {
	authHandler := JWTMiddleware(secret, handler)
	url = http.MethodPatch + " " + url
	router.HandleFunc(url, authHandler)
}

func Delete(router *http.ServeMux, url string, handler http.HandlerFunc) {
	url = http.MethodDelete + " " + url
	router.HandleFunc(url, handler)
}

func AuthDelete(router *http.ServeMux, url string, handler http.HandlerFunc, secret string) {
	authHandler := JWTMiddleware(secret, handler)
	url = http.MethodDelete + " " + url
	router.HandleFunc(url, authHandler)
}

func ChainMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
