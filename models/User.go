package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	IsActive bool   `json:"isActive"`
}
type UserCreate struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Image  string `json:"image"`
}
type UserUpdate struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (UserCreate) TableName() string { return "users" }
func (UserUpdate) TableName() string { return "users" }

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}

type ProductWithImgae struct {
	ID     int64
	Title  string
	Vendor string
	Image  string
	// Urls   Image
}

type Image struct {
	ProductID  int64   `json:"product_id"`
	ID         int64   `json:"id"`
	Position   int     `json:"position"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Alt        string  `json:"alt"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	Src        string  `json:"src"`
	VariantIDs []int64 `json:"variant_ids"`
	AdminAPIID string  `json:"admin_graphql_api_id"`
}

func (ProductWithImgae) TableName() string { return "products" }
