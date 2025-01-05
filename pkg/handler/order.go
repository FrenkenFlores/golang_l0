package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getOrder(c *gin.Context) {
	id := c.Param("id")
	err, message := h.service.GetOrder(id)
	if err != 200 {
		c.AbortWithStatusJSON(http.StatusBadRequest, message)
	}
	c.JSON(http.StatusOK, message)
}
