package handlers

import (
	"Youtube/NicJackson/Microservices2/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle POST.. product...")

	prod := h.Context().Value(KeyProduct{}).(data.Product)

	// p.l.Printf("Prod: %#v\n", &prod)
	data.AddProduct(&prod)
}

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Not able to marshal json..", http.StatusInternalServerError)
	}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle PUT.. product...")

	vars := mux.Vars(h)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := h.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product still not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {

		prod := data.Product{}
		err := prod.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Not able to unmarshal json..", http.StatusBadRequest)
			return
		}

		// ctx := h.Context().WithValue(KeyProduct{}, prod)
		ctx := context.WithValue(h.Context(), KeyProduct{}, prod)
		req := h.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
