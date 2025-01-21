package repository

type OrderDetailRepositoryInterface interface {
}

type OrderDetailRepository struct {
	BaseRepository *BaseRepository
}

func OrderDetailRepositoryInit(baseRepo *BaseRepository) *OrderDetailRepository {
	return &OrderDetailRepository{
		BaseRepository: baseRepo,
	}
}
