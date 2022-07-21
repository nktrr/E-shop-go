package handler

import (
	"E-shop-go/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllProducts(c *gin.Context) {
	page, err1 := strconv.Atoi(c.Query("page"))
	size, err2 := strconv.Atoi(c.Query("size"))
	if err1 != nil || err2 != nil {
		newErrorResponse(c, http.StatusBadRequest, err1.Error())
	}
	products, err := h.services.ProductInfo.GetAll(page, size)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: products,
	})
}

type getAllListsResponse struct {
	Data []entity.ProductInfo `json:"data"`
}
