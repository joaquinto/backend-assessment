package domain

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	PhoneNumber      string `json:"phoneNumber"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Region     string `json:"region"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserSuccessResponse struct {
	BaseResponse
	Data User `json:"data"`
}

type UsersSuccessResponse struct {
	BaseResponse
	TotalRecord int `json:"totalRecord"`
	Data []User `json:"data"`
}
