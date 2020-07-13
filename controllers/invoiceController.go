package controllers

import(
	"fmt"
	"strconv"	
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	"github.com/myrachanto/astore/service"
)
 
var (
	InvoiceController invoiceController = invoiceController{}
)
type invoiceController struct{ }
/////////controllers/////////////////
func (controller invoiceController) Create(c echo.Context) error {
	invoice := &model.Invoice{}
	
	if err := c.Bind(invoice); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	createdinvoice, err1 := service.Invoiceservice.Create(invoice)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdinvoice)
}
func (controller invoiceController) GetAll(c echo.Context) error {
	invoices := []model.Invoice{}
	invoices, err3 := service.Invoiceservice.GetAll(invoices)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, invoices)
} 
func (controller invoiceController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	invoice, problem := service.Invoiceservice.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, invoice)	
}

func (controller invoiceController) Update(c echo.Context) error {
		
	invoice :=  &model.Invoice{}
	if err := c.Bind(invoice); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedinvoice, problem := service.Invoiceservice.Update(id, invoice)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedinvoice)
}

func (controller invoiceController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.Invoiceservice.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}
