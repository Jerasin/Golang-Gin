package repository

type ProductRepositoryInterface interface {
}

type ProductRepository struct {
	BaseRepository *BaseRepository
}

func ProductRepositoryInit(baseRepo *BaseRepository) *ProductRepository {
	return &ProductRepository{
		BaseRepository: baseRepo,
	}
}
