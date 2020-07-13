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
	TransactionController transactionController = transactionController{}
)
type transactionController struct{ }
/////////controllers/////////////////
func (controller transactionController) Create(c echo.Context) error {
	transaction := &model.Transaction{}
	
	if err := c.Bind(transaction); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	createdtransaction, err1 := service.Transactionservice.Create(transaction)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdtransaction)
}
func (controller transactionController) GetAll(c echo.Context) error {
	transactions := []model.Transaction{}
	transactions, err3 := service.Transactionservice.GetAll(transactions)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, transactions)
} 
func (controller transactionController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	transaction, problem := service.Transactionservice.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, transaction)	
}

func (controller transactionController) Update(c echo.Context) error {
		
	transaction :=  &model.Transaction{}
	if err := c.Bind(transaction); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedtransaction, problem := service.Transactionservice.Update(id, transaction)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedtransaction)
}

func (controller transactionController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.Transactionservice.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}
