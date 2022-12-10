package handlers

import (
	"Youtube/NicJackson/Microservices/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {

	//GET
	if h.Method == http.MethodGet {
		p.getProducts(rw, h)
		return
	}

	//POST
	if h.Method == http.MethodPost {
		p.addProduct(rw, h)
		return
	}

	//PUT
	if h.Method == http.MethodPut {
		// first get the id to update
		// p.l.Println("PUT", h.URL.Path)
		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(h.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI... more than one id", g)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture.....")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number: ", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		// p.l.Println("Id is ", id)
		p.updateProduct(id, rw, h)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) addProduct(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle POST.. product...")

	prod := &data.Product{}
	err := prod.FromJSON(h.Body)
	if err != nil {
		http.Error(rw, "Not able to unmarshal json..", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()

	// we will convert the lp to json.
	// marshal allocates memory and it is slow.. so we will use json.Encoder.. it streams the data and also it is fast..
	// d, err := json.Marshal(lp)

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Not able to marshal json..", http.StatusInternalServerError)
	}
	// rw.Write(err)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle PUT.. product...")

	prod := &data.Product{}
	err := prod.FromJSON(h.Body)
	if err != nil {
		http.Error(rw, "Not able to unmarshal json..", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v\n", prod)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product still not found", http.StatusInternalServerError)
		return
	}

}
