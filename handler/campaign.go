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

	response := helper.APIResponse("Get Campaign success", 200, "success", campaigns)
	c.JSON(200, response)
}
