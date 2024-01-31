package dto

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateUserINput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
