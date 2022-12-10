package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// marshal allocates memory and it is slow.. so we will use json.Encoder.. it streams the data and also it is fast..
func (p *Products) ToJSON(w io.Writer) error {
	en := json.NewEncoder(w)
	return en.Encode(p)
}

// unmarshal
func (p *Product) FromJSON(r io.Reader) error {
	de := json.NewDecoder(r)
	return de.Decode(p)
}

func UpdateProduct(id int, p *Product) error {

	_, pos, err := findProduct(id)

	if err != nil {
		// log.Fatal(err)
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found!!!")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Frothy Coffee",
		Price:       3.25,
		SKU:         "hij432",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short strong coffee (no milk)",
		Price:       1.55,
		SKU:         "xys121",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
