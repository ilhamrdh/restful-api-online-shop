package products

type CreateProductRequest struct {
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

type ListProductRequest struct {
	Cursor int `query:"cursor" json:"cursor"`
	Size   int `query:"size" json:"size"`
}

func (l ListProductRequest) GenerateDefaultValue() ListProductRequest {
	if l.Cursor < 0 {
		l.Cursor = 0
	}
	if l.Size <= 0 {
		l.Size = 10
	}
	return l
}
