package controllers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/config"
	customerRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/customers"
	"github.com/ortizdavid/go-nopain/httputils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerController struct {
	customerRepository customerRepo.CustomerRepository
	logger *zap.Logger
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		customerRepository: *customerRepo.NewCustomerRepository(db),
		logger:             config.NewLogger("customners.log"),
	}
}

func (c *CustomerController) RegisterRoutes(router *http.ServeMux, db *gorm.DB) {
	router.HandleFunc("POST /api/customers", c.createCustomer)
	router.HandleFunc("PUT /api/customers/{id}", c.updateCustomer)
	router.HandleFunc("GET /api/customers", c.getAllCustomers)
	router.HandleFunc("GET /api/customers/{id}", c.getCustomerById)
}


func (c * CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	
}


func (c * CustomerController) updateCustomer(w http.ResponseWriter, r *http.Request) {
	
}


func (c * CustomerController) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	currentPage, limit := GetCurrentPageAndLimit(r)
	count := int(c.customerRepository.Count())

	if count == 0 {
		httputils.WriteJsonError(w, "no customers found", http.StatusNotFound)
		return
	}
	customers, err := c.customerRepository.GetAll(currentPage, limit)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJsonPaginated(w, r, customers, count, currentPage, limit)
}


func (c * CustomerController) getCustomerById(w http.ResponseWriter, r *http.Request) {
	
}