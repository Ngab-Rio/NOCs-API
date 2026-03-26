package dto

type CreatePackageRequest struct {
	Name  string  `json:"name" validate:"required"`
	Speed int     `json:"speed" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type UpdatePackageRequest struct {
	Name   *string  `json:"name" validate:"omitempty"`
	Speed  *int     `json:"speed" validate:"omitempty"`
	Price  *float64 `json:"price" validate:"omitempty"`
	Status *string  `json:"status" validate:"omitempty"`
}

type PackageResponse struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Speed  int     `json:"speed"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}
