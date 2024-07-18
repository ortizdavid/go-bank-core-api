package repositories

import "context"

type IRepository[T any] interface {
	Create(ctx context.Context, entity T) error
	CreateBatch(ctx context.Context, entities []T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, entity T) error
	GetAll(ctx context.Context, limit int, offset int) ([]T, error)
	GetById(ctx context.Context, id int) (T, error)
	GetByUniqueId(ctx context.Context, uniqueId string) (T, error)
	ExistsRecord(ctx context.Context, field string, value string) (bool, error)
}
