package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"privileges-management/database"
	"privileges-management/model/dto"
	"privileges-management/server/repository"
	"privileges-management/shares"
	"strconv"
)

type ApprovalsHandler interface {
	Approve(*gin.Context)
	Deny(*gin.Context)
}

type ApprovalsHandlerImpl struct {
	requestsRepository repository.AccessRequestsRepository
	resourceRepository repository.ResourcesRepository
}

func NewApprovalsHandler(requestsRepo repository.AccessRequestsRepository, resourcesRepo repository.ResourcesRepository) ApprovalsHandler {
	return &ApprovalsHandlerImpl{requestsRepository: requestsRepo, resourceRepository: resourcesRepo}
}

func (ah *ApprovalsHandlerImpl) Approve(ctx *gin.Context) {
	var approval dto.ApprovalPayload
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&approval)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	storageName := database.GetBucketName(id)
	database.UpdateToBucket(storageName, strconv.Itoa(rand.Int()), approval.Share)

	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	requiredShares := ah.resourceRepository.Read(uint(intId)).MinSharesRequired
	providedShares, err := database.GetNumberOfItemsFromBucket(storageName)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	if providedShares >= requiredShares {
		log.Printf("SEKRETY ZDEKODOWANE")
		retrievedShares, err := database.GetAllItemsFromBucket(storageName)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		fmt.Printf("%v\n", retrievedShares)

		passes, err := shares.RetrieveCredentialsFromSecrets(retrievedShares)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		fmt.Printf("%v\n", passes)
	}

}

func (ah *ApprovalsHandlerImpl) Deny(ctx *gin.Context) {
	id := ctx.Param("id")
	storageName := database.GetBucketName(id)
	database.RemoveBucket(storageName)
}
