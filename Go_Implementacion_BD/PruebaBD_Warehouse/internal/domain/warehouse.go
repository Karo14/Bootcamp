package domain

type Warehouse struct {
    ID    int               `json:"id"`
    Name  string            `json:"nombre"`
    Address  int            `json:"address"`
    Telephone string        `json:"telephone"`
    Capacity string     	`json:"capacity"`
}

type WarehouseReport struct {
	Warehouse
	ProductCount int
}