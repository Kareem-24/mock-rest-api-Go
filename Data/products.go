package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID         int `json:"id"`
	Name       string
	Desc       string
	Price      float32
	SKU        string
	Createddon string `josn:"-"`
	Updatedon  string `josn:"-"`
	Deletedon  string `josn:"-"`
}
type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts() Products {
	return productList
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found ")

func FindProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}

	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	{
		ID:         1,
		Name:       "Latte",
		Desc:       "milky coffe",
		Price:      2.45,
		Createddon: time.Now().UTC().String(),
		Updatedon:  time.Now().UTC().String(),
	},
	{
		ID:         2,
		Name:       "Espresso",
		Desc:       "strong coffe , no milk",
		Price:      1.99,
		Createddon: time.Now().UTC().String(),
		Updatedon:  time.Now().UTC().String(),
	},
}
