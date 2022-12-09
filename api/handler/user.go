package handler

import (
	"log"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/storage"
)

func (h *HandlerV1) GetList(c *gin.Context) {

	users, err := storage.GetList(h.db)
	if err != nil {
		log.Fatalf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list"))
		return
	}

	c.JSON(http.StatusOK, users)
}
