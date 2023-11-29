package models

type Merchandise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewMerchandise(id, name string) *Merchandise {
	return &Merchandise{
		ID:   id,
		Name: name,
	}
}
