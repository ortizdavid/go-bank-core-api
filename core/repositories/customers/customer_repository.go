package repositories

import (
	entities "github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (repo *CustomerRepository) Create(customer entities.Customer) error {
	result := repo.db.Create(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CustomerRepository) CreateBatch(customers []entities.Customer) error {
	tx := repo.db.Begin()
	result := repo.db.Create(&customers)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (repo *CustomerRepository) Update(customer entities.Customer) error {
	result := repo.db.Save(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CustomerRepository) Delete(customer entities.Customer) error {
	result := repo.db.Delete(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CustomerRepository) GetAll(limit int, offest int) ([]entities.Customer, error) {
	var customers []entities.Customer
	result := repo.db.Find(&customers).Limit(limit).Offset(offest)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (repo *CustomerRepository) GetById(userId int64) (entities.Customer, error) {
	var customer entities.Customer
	result := repo.db.First(&customer, userId)
	if result.Error != nil {
		return entities.Customer{}, result.Error
	}
	return customer, nil
}

func (repo *CustomerRepository) GetByUniqueId(uniqueId string) (entities.Customer, error) {
	var customer entities.Customer
	result := repo.db.First(&customer, "unique_id", uniqueId)
	if result.Error != nil {
		return entities.Customer{}, result.Error
	}
	return customer, nil
}

func (repo *CustomerRepository) Count() int64 {
	var count int64
	result := repo.db.Count(&count)
	if result.Error != nil {
		return 0
	}
	return count
}
