package handler

import (
	"net/http"
	"startup_funding/campaign"
	"startup_funding/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap paramter
// tangkap handler ke service
// service yang menetukan repository mana yang di call

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
		response := helper.APIResponse("List of campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

	}
	response := helper.APIResponse("List of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

//
