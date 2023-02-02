package handler

import (
	"growfunding/campaign"
	"growfunding/helper"
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
