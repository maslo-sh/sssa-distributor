package handlers

import (
	"github.com/gin-gonic/gin"
	"privileges-management/database"
	"privileges-management/server/repository"
)

type ApprovalsHandler interface {
	Approve(*gin.Context)
	Deny(*gin.Context)
}

type ApprovalsHandlerImpl struct {
	requestsRepository repository.AccessRequestsRepository
}

func NewApprovalsHandler(requestsRepo repository.AccessRequestsRepository) ApprovalsHandler {
	return &ApprovalsHandlerImpl{requestsRepository: requestsRepo}
}

func (ah *ApprovalsHandlerImpl) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	storageName := database.GetBucketName(id)
	database.CreateBucketIfNotExists(storageName)
}

func (ah *ApprovalsHandlerImpl) Deny(ctx *gin.Context) {
	id := ctx.Param("id")
	storageName := database.GetBucketName(id)
	database.RemoveBucket(storageName)
}
