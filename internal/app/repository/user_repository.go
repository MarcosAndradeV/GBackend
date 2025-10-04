package repository

import (
	"context"
	"gbackend/internal/app/domain"
	"gorm.io/gorm"
)

type exampleUserRepository struct {
	db *gorm.DB
	id string
}

func NewUserRepository(db *gorm.DB, id string) domain.UserRepository {
	return &exampleUserRepository{
		db: db,
		id: id,
	}
}

func (ur *exampleUserRepository) Create(c context.Context, user *domain.ExampleUser) error {
	ctx := context.Background()
	return gorm.G[domain.ExampleUser](ur.db).Create(ctx, user)
}

func (ur *exampleUserRepository) Fetch(c context.Context) ([]domain.ExampleUser, error) {
	ctx := context.Background()
	return gorm.G[domain.ExampleUser](ur.db).Find(ctx)
}
