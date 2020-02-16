package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wicoady1/demo/validator"
)

type Middleware func(httprouter.Handle) httprouter.Handle

func MultipleMiddlewares(h httprouter.Handle, m ...Middleware) httprouter.Handle {
	if len(m) < 1 {
		return h
	}
	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func CustomerIDValidatorMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		customerValid := validator.ValidCustomerID(r)
		if !customerValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println("Customer is valid!")
		next(w, r, ps)
	}
}

func LogMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//logging functions
		log.Println("Log Middleware")
		next(w, r, ps)
	}
}

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//auth functions
		log.Println("Auth middleware")
		next(w, r, ps)
	}
}
