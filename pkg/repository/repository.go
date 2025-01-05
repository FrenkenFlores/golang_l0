package repository

type Repository struct {
	Order
}

type Order interface {
	GetOrder(id string) (int, map[string]any)
}

func NewRepository(cfg Config) (*Repository, error) {
	order, err := NewPostgresDb(cfg)
	return &Repository{
		Order: order,
	}, err
}
