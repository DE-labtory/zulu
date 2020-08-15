package api

import "github.com/gin-gonic/gin"

func badrequestError(context *gin.Context, err error) {
	context.JSON(400, gin.H{"error": err.Error()})
}

func internalServerError(context *gin.Context, err error) {
	context.JSON(500, gin.H{"error": err.Error()})
}
