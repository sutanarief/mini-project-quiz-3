package structs

type Category struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Book struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_url    string `json:"image_url"`
	Release_year int64  `json:"release_year"`
	Price        string `json:"price"`
	Total_page   int64  `json:"total_page"`
	Thickness    string `json:"thickness"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
	Category_id  int64  `json:"category_id"`
}
