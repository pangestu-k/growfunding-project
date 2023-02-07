package handler

import (
	"growfunding/helper"
	"growfunding/transaction"
	"growfunding/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetTransactionCampaignInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Get data Transaction failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)

	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("Get data Transaction failed", 400, "error", errorMessage)
		c.JSON(400, response)
		return
	}

	transactionResponse := transaction.TransactionsFormatter(transactions)
	response := helper.APIResponse("Get data Transaction success", 200, "success", transactionResponse)
	c.JSON(200, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionByUserID(userID)

	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("Get data Transaction failed", 400, "error", errorMessage)
		c.JSON(400, response)
		return
	}

	response := helper.APIResponse("Get data Transaction success", 200, "success", transaction.FormatUsersTransaction(transactions))
	c.JSON(200, response)
}
