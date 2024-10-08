package handler

import ( 
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func sendError(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}){
	ctx.Header("Context-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreatingResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}