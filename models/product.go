package models

type Product struct {
	ID              int              `json:"id"`
	SKU             int              `json:"sku"`
	Imagen          string           `json:"imagen"`
	Nombre          string           `json:"nombre"`
	Descripcion     string           `json:"descripcion"`
	Caracteristicas []Characteristic `json:"caracteristicas"`
	Marca           string           `json:"marca"`
	Precio          int              `json:"precio"`
}
