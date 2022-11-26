package core

type Store struct {
	StoreId        int    `json:"store_id"`
	Name           string `json:"name"`
	StoreAddressId *int   `json:"store_address_id,omitempty"`
	SupplierId     *int   `json:"supplier_id,omitempty"`
}

type CreateStore struct {
	Name           string `json:"name"`
	StoreAddressId *int   `json:"store_address_id,omitempty"`
	SupplierId     *int   `json:"supplier_id,omitempty"`
}

type CreateStoreDTO struct {
	Name           string            `json:"name"`
	Creator        CreateEmployeeDTO `json:"creator,omitempty"`
	StoreAddressId *int              `json:"store_address_id,omitempty"`
	SupplierId     *int              `json:"supplier_id,omitempty"`
}

func NewCreateStoreFromDTO(dto *CreateStoreDTO, addressId *int, supplierId *int) *CreateStore {
	return &CreateStore{
		Name:           dto.Name,
		StoreAddressId: addressId,
		SupplierId:     supplierId,
	}
}
