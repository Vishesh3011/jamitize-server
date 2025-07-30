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

func MethodGuard(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func Post(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	mux.HandleFunc(path, MethodGuard(http.MethodPost, handler))
}

func Get(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	mux.HandleFunc(path, MethodGuard(http.MethodGet, handler))
}
