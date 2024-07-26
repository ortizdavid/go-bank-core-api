package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ortizdavid/go-bank-core-api/common/apperrors"
	"github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"github.com/ortizdavid/go-bank-core-api/core/repositories/customers"
	"github.com/ortizdavid/go-bank-core-api/common/helpers"
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


func (s *CustomerService) UpdateCustomerContacts(r *http.Request, ctx context.Context, request entities.UpdateCustomerContactRequest) error {
	if err := serialization.DecodeJson(r.Body, &request); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	customer, err := s.repository.GetById(ctx, request.CustomerId)
	if err != nil {
		return apperrors.NewNotFoundError("Customer not found")
	}
	customer.Email = request.Email
	customer.Phone = request.Phone
	s.repository.Update(ctx, customer)
	return nil
}


func (s *CustomerService) ChangeCustomerType(r *http.Request, ctx context.Context, request entities.ChangeCustomerTypeRequest) error {
	if err := serialization.DecodeJson(r.Body, &request); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	customer, err := s.repository.GetById(ctx, request.CustomerId)
	if err != nil {
		return apperrors.NewNotFoundError("Customer not found")
	}
	customer.CustomerType = request.NewType
	s.repository.Update(ctx, customer)
	return nil
}


func (s *CustomerService) ChangeCustomerStatus(r *http.Request, ctx context.Context, request entities.ChangeCustomerStatusRequest) error {
	if err := serialization.DecodeJson(r.Body, &request); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	customer, err := s.repository.GetById(ctx, request.CustomerId)
	if err != nil {
		return apperrors.NewNotFoundError("Customer not found")
	}
	customer.CustomerStatus = request.NewStatus
	s.repository.Update(ctx, customer)
	return nil
}


func (s *CustomerService) DeleteCustomer(ctx context.Context, customerId int64) error {
	customer, err := s.repository.GetById(ctx, customerId)
	if err != nil {
		return apperrors.NewNotFoundError("Customer not found")
	}
	s.repository.Delete(ctx, customer)
	return nil
}


func (s *CustomerService) GetAllCustomers(r *http.Request, ctx context.Context, params helpers.PaginationParams) (*httputils.Pagination[entities.CustomerData], error) {
	if params.CurrentPage < 0  || params.Limit < 1 {
		return nil, apperrors.NewBadRequestError("Incorrect current page or limit")
	}
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, apperrors.NewNotFoundError("No customers found")
	}
	customers, err := s.repository.GetAll(ctx, params.Limit, params.CurrentPage)
	if err != nil {
		return nil, err
	}
	pagination, err := httputils.NewPagination(r, customers, count, params.CurrentPage, params.Limit)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}


func (s *CustomerService) GetCustomerById(ctx context.Context, customerId int64) (entities.CustomerData, error) {
	if customerId < 0 {
		return entities.CustomerData{}, apperrors.NewBadRequestError("CustomerId must be greather than 0")
	}
	customer, err := s.repository.GetDataById(ctx, customerId)
	if err != nil {
		return entities.CustomerData{}, apperrors.NewNotFoundError("Customer not found")
	}
	return customer, nil
}


func (s *CustomerService) GetCustomerByUniqueId(ctx context.Context, uniqueId string) (entities.CustomerData, error) {
	if uniqueId == "" {
		return entities.CustomerData{}, apperrors.NewBadRequestError("UniqueId cannot be null")
	}
	customer, err := s.repository.GetByDataUniqueId(ctx, uniqueId)
	if err != nil {
		return entities.CustomerData{}, apperrors.NewNotFoundError("Customer not found")
	}
	return customer, nil
}


func (s *CustomerService) checkDuplicatedValues(ctx context.Context, field string, value string) error {
	exists, err := s.repository.ExistsRecord(ctx, field, value)
	if err != nil {
	  return fmt.Errorf("error checking %s existence: %w", field, err)
	}
	if exists {
	  return fmt.Errorf("%s '%s' already exists", field, value)
	}
	return nil
}