package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wicoady1/demo/service"
	"github.com/wicoady1/demo/validator"
)

func GetCart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerValid := validator.ValidCustomerID(r)
	if !customerValid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	resp := service.GetCart(r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func InitCart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerValid := validator.ValidCustomerID(r)
	if !customerValid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	resp := service.InitCart(r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func AddItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerValid := validator.ValidCustomerID(r)
	if !customerValid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	resp := service.AddItems(r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func DeleteItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerValid := validator.ValidCustomerID(r)
	if !customerValid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	resp := service.DeleteItems(r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
