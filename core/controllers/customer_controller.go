package controllers

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-bank-core-api/config"
	entities "github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"github.com/ortizdavid/go-bank-core-api/core/services/customers"
	"github.com/ortizdavid/go-bank-core-api/helpers"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerController struct {
	service services.CustomerService
	infoLogger *zap.Logger
	errorLogger *zap.Logger
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		service: *services.NewCustomerService(db),
		infoLogger:  config.NewLogger("customers-info.log"),
		errorLogger:  config.NewLogger("customers-error.log"),
	}
}

func (c *CustomerController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/customers", c.getAllCustomers)
	router.HandleFunc("GET /api/customers/{id}", c.getCustomerById)
	router.HandleFunc("GET /api/customers/by-uuid/{unique_id}", c.getCustomerByUniqueId)
	router.HandleFunc("POST /api/customers", c.createCustomer)
	router.HandleFunc("PUT /api/customers/change-type", c.changeCustomerType)
	router.HandleFunc("PUT /api/customers/change-status", c.changeCustomerStatus)
	router.HandleFunc("PUT /api/customers/update-contacts", c.updateCustomerContacts)
	router.HandleFunc("DELETE /api/customers/{id}", c.deleteCustomer)
	
}

func (c * CustomerController) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := helpers.GetPaginationParams(r)
	customers, err := c.service.GetAllCustomers(r, r.Context(), params)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, customers)
}

func (c * CustomerController) getCustomerById(w http.ResponseWriter, r *http.Request) {
	id := conversion.StringToInt64(r.PathValue("id"))
	customer, err := c.service.GetCustomerById(r.Context(), id)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (c * CustomerController) getCustomerByUniqueId(w http.ResponseWriter, r *http.Request) {
	uniqueId := r.PathValue("unique_id")
	customer, err := c.service.GetCustomerByUniqueId(r.Context(), uniqueId)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (c * CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	var request entities.CreateCustomerRequest
	err := c.service.CreateCustomer(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		c.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%s' created", request.CustomerName)
	c.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (c * CustomerController) changeCustomerType(w http.ResponseWriter, r *http.Request) {
	var request entities.ChangeCustomerTypeRequest
	err := c.service.ChangeCustomerType(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		c.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' Type Changed", request.CustomerId)
	c.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (c * CustomerController) changeCustomerStatus(w http.ResponseWriter, r *http.Request) {
	var request entities.ChangeCustomerStatusRequest
	err := c.service.ChangeCustomerStatus(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		c.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' Status Changed", request.CustomerId)
	c.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (c * CustomerController) updateCustomerContacts(w http.ResponseWriter, r *http.Request) {
	var request entities.UpdateCustomerContactRequest
	err := c.service.UpdateCustomerContacts(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		c.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' contacts updated", request.CustomerId)
	c.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (c * CustomerController) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := conversion.StringToInt64(r.PathValue("id"))
	err := c.service.DeleteCustomer(r.Context(), customerId)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		c.errorLogger.Error(err.Error())
		return 
	}
	msg := fmt.Sprintf("Customer '%d' deleted", customerId)
	c.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusNoContent, msg)
}


