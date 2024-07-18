package controllers

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-bank-core-api/config"
	entities "github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"github.com/ortizdavid/go-bank-core-api/core/services/customers"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerController struct {
	service services.CustomerService
	logger *zap.Logger
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		service: *services.NewCustomerService(db),
		logger:  config.NewLogger("customners.log"),
	}
}

func (c *CustomerController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/customers", c.getAllCustomers)
	router.HandleFunc("GET /api/customers/{id}", c.getCustomerById)
	router.HandleFunc("GET /api/customers/by-uuid/{unique_id}", c.getCustomerByUniqueId)
	router.HandleFunc("POST /api/customers", c.createCustomer)
	router.HandleFunc("PUT /api/customers/{id}", c.updateCustomer)
	router.HandleFunc("DELETE /api/customers/{id}", c.deleteCustomer)
	
}

func (c * CustomerController) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	currentPage, limit := getCurrentPageAndLimit(r)
	customers, err := c.service.GetAllCustomers(r, r.Context(), currentPage, limit)
	if err != nil {
		handleCustomErrors(w, err)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, customers)
}

func (c * CustomerController) getCustomerById(w http.ResponseWriter, r *http.Request) {
	id := conversion.StringToInt64(r.PathValue("id"))
	customer, err := c.service.GetCustomerById(r.Context(), id)
	if err != nil {
		handleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (c * CustomerController) getCustomerByUniqueId(w http.ResponseWriter, r *http.Request) {
	uniqueId := r.PathValue("unique_id")
	customer, err := c.service.GetCustomerByUniqueId(r.Context(), uniqueId)
	if err != nil {
		handleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (c * CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	var request entities.CreateCustomerRequest
	if err := c.service.CreateCustomer(r, r.Context(), request); err != nil {
		handleCustomErrors(w, err)
		c.logger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%s' created", request.CustomerName)
	c.logger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (c * CustomerController) updateCustomer(w http.ResponseWriter, r *http.Request) {
	
}

func (c * CustomerController) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	
}
