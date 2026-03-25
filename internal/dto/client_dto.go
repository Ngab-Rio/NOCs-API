package dto

type CreateClientRequest struct {
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Phone     string  `json:"phone" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Status    string  `json:"status" binding:"required,oneof=active inactive"`
}

type UpdateClientRequest struct {
	Name      *string  `json:"name"`
	Email     *string  `json:"email" binding:"omitempty,email"`
	Phone     *string  `json:"phone"`
	Address   *string  `json:"address"`
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
	Status    *string  `json:"status" binding:"omitempty,oneof=active inactive"`
}

type ClientResponse struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Status    string  `json:"status"`
}

type GetClientsRequest struct {
	Limit  int    `form:"limit" binding:"required,min=1,max=100"`
	Offset int    `form:"offset" binding:"min=0"`
	Search string `form:"search"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

type GetClientsResponse struct {
	Data   []ClientResponse `json:"data"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
	Total  int64            `json:"total"`
}
