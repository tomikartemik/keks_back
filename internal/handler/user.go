package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	userIDStr := c.Query("tg_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	user, err := h.services.GetUserById(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserAsSellerByID(c *gin.Context) {
	telegramIDStr := c.Query("tg_id")

	userAsSeller, err := h.services.GetUserAsSellerByID(telegramIDStr)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID: "+err.Error())
	}

	c.JSON(http.StatusOK, userAsSeller)
}
