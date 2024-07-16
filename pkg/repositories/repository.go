package repositories

type IRepository[T any] interface {
	Create(entity T) error
	CreateBatch(entities []T) error
	Update(entity T) error
	Delete(entity T) error
	GetAll() ([]T, error)
	GetById(id int) (T, error)
	GetByUniqueId(uniqueId string) (T, error)
	ExistsRecord(field string, value string) (bool, error)
}
