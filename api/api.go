package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.GET("/user", handlerV1.GetList)

}
