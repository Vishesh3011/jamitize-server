package routes

import (
	"encoding/json"
	"example/errors"
	"example/internal/core/application"
	"net/http"
	"runtime"
	"runtime/debug"
)

type requestHandler[T any] func(application.Application, *http.ResponseWriter, *http.Request) (*T, errors.AppError)

func HandleRequest[T any](application application.Application, handler requestHandler[T]) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
					json.NewEncoder(writer).Encode(appError.Json())
				} else {
					application.Logger().Error("API Request error: %v", handlerError)
					http.Error(writer, handlerError.Error(), http.StatusInternalServerError)
				}
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(map[string]any{
				"data": response,
			})
		}()
		response, handlerError = handler(application, &writer, request)
	}
}

func Post(router *http.ServeMux, url string, handler http.HandlerFunc) {
	router.HandleFunc(url, handler)
}

func Get(router *http.ServeMux, url string, handler http.HandlerFunc) {
	router.HandleFunc(url, handler)
}

func Put(router *http.ServeMux, url string, handler http.HandlerFunc) {
	router.HandleFunc(url, handler)
}

func Delete(router *http.ServeMux, url string, handler http.HandlerFunc) {
	router.HandleFunc(url, handler)
}

func ChainMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
