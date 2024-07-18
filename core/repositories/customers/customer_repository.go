package repositories

import (
	"context"
	"errors"

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

func (repo *CustomerRepository) Create(ctx context.Context, customer entities.Customer) error {
	result := repo.db.WithContext(ctx).Create(&customer)
	return result.Error
}

func (repo *CustomerRepository) CreateBatch(ctx context.Context, customers []entities.Customer) error {
	tx := repo.db.Begin()
	result := repo.db.WithContext(ctx).Create(&customers)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func (repo *CustomerRepository) Update(ctx context.Context, customer entities.Customer) error {
	result := repo.db.WithContext(ctx).Save(&customer)
	return result.Error
}

func (repo *CustomerRepository) Delete(ctx context.Context, customer entities.Customer) error {
	result := repo.db.WithContext(ctx).Delete(&customer)
	return result.Error
}

func (repo *CustomerRepository) GetAll(ctx context.Context, limit int, offest int) ([]entities.Customer, error) {
	var customers []entities.Customer
	result := repo.db.WithContext(ctx).Find(&customers).Limit(limit).Offset(offest)
	return customers, result.Error
}

func (repo *CustomerRepository) GetById(ctx context.Context, customerId int64) (entities.Customer, error) {
	var customer entities.Customer
	result := repo.db.WithContext(ctx).First(&customer, customerId)
	if result.Error != nil {
		return entities.Customer{}, result.Error
	}
	return customer, result.Error
}

func (repo *CustomerRepository) GetByUniqueId(ctx context.Context, uniqueId string) (entities.Customer, error) {
	var customer entities.Customer
	result := repo.db.WithContext(ctx).First(&customer, "unique_id", uniqueId)
	if result.Error != nil {
		return entities.Customer{}, result.Error
	}
	return customer, result.Error
}

func (repo *CustomerRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Count(&count)
	return count, result.Error
}

func (repo *CustomerRepository) ExistsRecord(ctx context.Context, field string, value string) (bool, error) {
    var count int64
    // Validate the field to avoid SQL injection
    validFields := map[string]bool{
        "identification_number":  true,
        "email":                  true,
        "phone":                  true,
    }
    if !validFields[field] {
        return false, errors.New("invalid field name")
    }
    // Construct and execute the query
    query := repo.db.WithContext(ctx).Table("customers").Where(field+" = ?", value).Count(&count)
    if query.Error != nil {
        return false, query.Error
    }
    return count > 0, nil
}


