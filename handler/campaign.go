package handler

import (
	"growfunding/campaign"
	"growfunding/helper"
	"growfunding/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Get Campaign Fail", 422, "error", nil)
		c.JSON(422, response)
		return
	}

	response := helper.APIResponse("Get Campaign success", 200, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(200, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Get Detail Campaign Fail", 400, "error", nil)
		c.JSON(400, response)
		return
	}

	campaignDetail, err := h.service.GetCampaign(input)
	if err != nil {
		response := helper.APIResponse("Get Detail Campaign Fail", 400, "error", nil)
		c.JSON(400, response)
		return
	}

	response := helper.APIResponse("Get Detail Campaign sucess", 200, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(200, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	input := campaign.CreateCampaignInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Create Campaign failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Create Campaign failed", 400, "error", nil)
		c.JSON(400, response)
		return
	}

	campaignResponse := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Create user Success", 200, "success", campaignResponse)
	c.JSON(200, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	inputData := campaign.CreateCampaignInput{}
	err := c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Create Campaign failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	inputID := campaign.GetCampaignDetailInput{}
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Create Campaign failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	update, err := h.service.UpdataeCampaign(inputID, inputData)

	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Update Campaign failed", 400, "error", errorMessage)
		c.JSON(400, response)
		return
	}

	campaignResponse := campaign.FormatCampaign(update)
	response := helper.APIResponse("Update Campaign success", 200, "success", campaignResponse)
	c.JSON(200, response)
}
