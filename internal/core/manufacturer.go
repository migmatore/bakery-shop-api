package core

type Manufacturer struct {
	ManufacturerId int    `json:"manufacturer_id"`
	Name           string `json:"name"`
	SupplierId     *int   `json:"supplier_id,omitempty"`
}
