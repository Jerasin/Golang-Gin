package repository

type ProductCategoryRepositoryInterface interface{}

type ProductCategoryRepository struct {
	BaseRepository *BaseRepository
}

func ProductCategoryRepositoryInit(baseRepo *BaseRepository) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		BaseRepository: baseRepo,
	}
}
