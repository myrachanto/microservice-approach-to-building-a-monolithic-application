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
	ProductController productController = productController{}
)
type productController struct{ }
/////////controllers/////////////////
func (controller productController) Create(c echo.Context) error {
	product := &model.Product{}
	
	if err := c.Bind(product); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	createdproduct, err1 := service.Productservice.Create(product)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdproduct)
}
func (controller productController) GetAll(c echo.Context) error {
	products := []model.Product{}
	products, err3 := service.Productservice.GetAll(products)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, products)
} 
func (controller productController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	product, problem := service.Productservice.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, product)	
}

func (controller productController) Update(c echo.Context) error {
		
	product :=  &model.Product{}
	if err := c.Bind(product); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedproduct, problem := service.Productservice.Update(id, product)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedproduct)
}

func (controller productController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.Productservice.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}
