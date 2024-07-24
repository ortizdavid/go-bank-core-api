package controllers

import (
	"fmt"
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/config"
	"github.com/ortizdavid/go-bank-core-api/common/helpers"
	entities "github.com/ortizdavid/go-bank-core-api/core/entities/customers"
	"github.com/ortizdavid/go-bank-core-api/core/services/customers"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerController struct {
	service *services.CustomerService
	infoLogger *zap.Logger
	errorLogger *zap.Logger
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		service: services.NewCustomerService(db),
		infoLogger:  config.NewLogger("customers-info.log"),
		errorLogger:  config.NewLogger("customers-error.log"),
	}
}

func (ctrl *CustomerController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/customers", ctrl.getAllCustomers)
	router.HandleFunc("GET /api/customers/{id}", ctrl.getCustomerById)
	router.HandleFunc("GET /api/customers/by-uuid/{unique_id}", ctrl.getCustomerByUniqueId)
	router.HandleFunc("POST /api/customers", ctrl.createCustomer)
	router.HandleFunc("PUT /api/customers/change-type", ctrl.changeCustomerType)
	router.HandleFunc("PUT /api/customers/change-status", ctrl.changeCustomerStatus)
	router.HandleFunc("PUT /api/customers/update-contacts", ctrl.updateCustomerContacts)
	router.HandleFunc("DELETE /api/customers/{id}", ctrl.deleteCustomer)
}

func (ctrl *CustomerController) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	params := helpers.GetPaginationParams(r)
	customers, err := ctrl.service.GetAllCustomers(r, r.Context(), params)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJsonSimple(w, http.StatusOK, customers)
}

func (ctrl *CustomerController) getCustomerById(w http.ResponseWriter, r *http.Request) {
	id := conversion.StringToInt64(r.PathValue("id"))
	customer, err := ctrl.service.GetCustomerById(r.Context(), id)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (ctrl *CustomerController) getCustomerByUniqueId(w http.ResponseWriter, r *http.Request) {
	uniqueId := r.PathValue("unique_id")
	customer, err := ctrl.service.GetCustomerByUniqueId(r.Context(), uniqueId)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		return
	}
	httputils.WriteJson(w, http.StatusOK, customer)
}

func (ctrl *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	var request entities.CreateCustomerRequest
	err := ctrl.service.CreateCustomer(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		ctrl.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%s' created", request.CustomerName)
	ctrl.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (ctrl *CustomerController) changeCustomerType(w http.ResponseWriter, r *http.Request) {
	var request entities.ChangeCustomerTypeRequest
	err := ctrl.service.ChangeCustomerType(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		ctrl.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' type changed", request.CustomerId)
	ctrl.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (ctrl *CustomerController) changeCustomerStatus(w http.ResponseWriter, r *http.Request) {
	var request entities.ChangeCustomerStatusRequest
	err := ctrl.service.ChangeCustomerStatus(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		ctrl.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' status changed", request.CustomerId)
	ctrl.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (ctrl *CustomerController) updateCustomerContacts(w http.ResponseWriter, r *http.Request) {
	var request entities.UpdateCustomerContactRequest
	err := ctrl.service.UpdateCustomerContacts(r, r.Context(), request)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		ctrl.errorLogger.Error(err.Error())
		return
	}
	msg := fmt.Sprintf("Customer '%d' contacts updated", request.CustomerId)
	ctrl.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusCreated, msg)
}

func (ctrl *CustomerController) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := conversion.StringToInt64(r.PathValue("id"))
	err := ctrl.service.DeleteCustomer(r.Context(), customerId)
	if err != nil {
		helpers.HandleCustomErrors(w, err)
		ctrl.errorLogger.Error(err.Error())
		return 
	}
	msg := fmt.Sprintf("Customer '%d' deleted", customerId)
	ctrl.infoLogger.Info(msg)
	httputils.WriteJson(w, http.StatusNoContent, msg)
}


