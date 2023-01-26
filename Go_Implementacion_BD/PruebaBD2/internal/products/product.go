package products

import "time"

type Product struct {
    ID    int               `json:"id"`
    Name  string            `json:"nombre"`
    Quantity  int           `json:"quantity"`
    Code_value string       `json:"code_value"`
    Is_published string     `json:"is_published"`
    Expiration time.Time    `json:"expiration"`
    Price float64           `json:"price"`
}