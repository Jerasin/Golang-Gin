package repository

type OrderRepositoryInterface interface {
}

type OrderRepository struct {
	BaseRepository *BaseRepository
}

func OrderRepositoryInit(baseRepo *BaseRepository) *OrderRepository {
	return &OrderRepository{
		BaseRepository: baseRepo,
	}
}
