package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ortizdavid/go-bank-core-api/apperrors"
	"github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"github.com/ortizdavid/go-bank-core-api/core/repositories/customers"
	"github.com/ortizdavid/go-nopain/datetime"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"gorm.io/gorm"
)

type CustomerService struct {
	repository repositories.CustomerRepository
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		repository: *repositories.NewCustomerRepository(db),
	}
}

func (s *CustomerService) CreateCustomer(r *http.Request, ctx context.Context, request entities.CreateCustomerRequest) error {
	if err := serialization.DecodeJson(r.Body, &request); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	if err := s.checkDuplicatedValues(ctx, "identification_number", request.IdentificationNumber); err != nil {
		return apperrors.NewConflictError(err.Error())
	}
	if err := s.checkDuplicatedValues(ctx, "email", request.Email); err != nil {
		return apperrors.NewConflictError(err.Error())
	}
	if err := s.checkDuplicatedValues(ctx, "phone", request.Phone); err != nil {
		return apperrors.NewConflictError(err.Error())
	}
	customer := entities.Customer{
		CustomerType: request.CustomerType,
		CustomerStatus: entities.CustomerStatusPending,
		CustomerName: request.CustomerName,
		IdentificationNumber: request.IdentificationNumber,
		Gender: request.Gender,
		BirtDate: datetime.StringToDate(request.BirthDate),
		Email: request.Email,
		Phone: request.Phone,
		Address: request.Address,
		UniqueId: encryption.GenerateUUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	return s.repository.Create(ctx, customer)
}

func (s *CustomerService) GetAllCustomers(r *http.Request, ctx context.Context, currentPage int, limit int) (*httputils.Pagination[entities.Customer], error) {
	if currentPage < 0  || limit < 1 {
		return nil, apperrors.NewBadRequestError("Incorrect current page or limit")
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No customers found")
	}
	customers, err := s.repository.GetAll(ctx, limit, currentPage)
	if err != nil {
		return nil, err
	}
	pagination, err := httputils.NewPagination(r, customers, int(count), currentPage, limit)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func (s *CustomerService) GetCustomerById(ctx context.Context, customerId int64) (entities.Customer, error) {
	if customerId < 0 {
		return entities.Customer{}, apperrors.NewBadRequestError("CustomerId must be greather than 0")
	}
	customer, err := s.repository.GetById(ctx, customerId)
	if err != nil {
		return entities.Customer{}, apperrors.NewNotFoundError("Customer not found")
	}
	return customer, nil
}

func (s *CustomerService) GetCustomerByUniqueId(ctx context.Context, uniqueId string) (entities.Customer, error) {
	if uniqueId == "" {
		return entities.Customer{}, apperrors.NewBadRequestError("UniqueId cannot be null")
	}
	customer, err := s.repository.GetByUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.Customer{}, apperrors.NewNotFoundError("Customer not found")
	}
	return customer, nil
}


func (s *CustomerService) checkDuplicatedValues(ctx context.Context, field string, value string) error {
	exists, err := s.repository.ExistsRecord(ctx, field, value)
	if err != nil {
	  return fmt.Errorf("error checking %s existence: %w", field, err)
	}
	if exists {
	  return apperrors.NewConflictError(fmt.Sprintf("%s '%s' already exists", field, value))
	}
	return nil
}