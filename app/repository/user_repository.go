package repository

type UserRepositoryInterface interface{}

type UserRepository struct {
	BaseRepository *BaseRepository
}

func UserRepositoryInit(baseRepo *BaseRepository) *UserRepository {
	return &UserRepository{
		BaseRepository: baseRepo,
	}
}
