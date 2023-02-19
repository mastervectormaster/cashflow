package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/api/v1/models"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/api/v1/responses"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "cash")
var validate = validator.New()

// Report Income             godoc
// @Summary      Submit New Income Report
// @Description  Takes a Income Report JSON and store in DB. Return saved JSON.
// @Tags         Cash
// @Produce      json
// @Param        book  body      models.Cash  true  "Book JSON"
// @Success      200   {object}  models.Cash
// @Router       /income [post]
func ReportIncome(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cash models.Cash
	defer cancel()

	//validate the request body
	if err := c.BindJSON(&cash); err != nil {
		c.JSON(http.StatusBadRequest, responses.CashResponse{Status: http.StatusBadRequest, Message: "bind error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&cash); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.CashResponse{Status: http.StatusBadRequest, Message: "validation error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newUser := models.Cash{
		Name:     cash.Name,
		Amount:   cash.Amount,
		Platform: cash.Platform,
		Url:      cash.Url,
		// Img:      cash.Img,
		Notes: cash.Notes,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.CashResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.CashResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"insertedID": result.InsertedID}})
}
