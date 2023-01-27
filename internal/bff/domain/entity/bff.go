package entity

type (
	Store struct {
		ID    uint32 `json:"id"`
		Name  string `json:"name"`
		Books []Book `json:"books,omitempty"`
	}

	Book struct {
		ID      uint32 `json:"id"`
		StoreID uint32 `json:"store_id"`
		Name    string `json:"name"`
	}
)
