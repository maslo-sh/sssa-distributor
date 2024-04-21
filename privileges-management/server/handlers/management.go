package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"privileges-management/model"
	"privileges-management/model/dto"
	"privileges-management/server/repository"
)

type ManagementHandler interface {
	AssignApproversToResource(*gin.Context)
	RegisterResource(ctx *gin.Context)
}

type ManagementHandlerImpl struct {
	permissionsRepository repository.ApprovingPermissionsRepository
	resourcesRepository   repository.ResourcesRepository
}

func NewManagementHandler(permissionsRepo repository.ApprovingPermissionsRepository, resourcesRepo repository.ResourcesRepository) ManagementHandler {
	return &ManagementHandlerImpl{permissionsRepository: permissionsRepo, resourcesRepository: resourcesRepo}
}

func (mh *ManagementHandlerImpl) AssignApproversToResource(ctx *gin.Context) {
	var req dto.ApproverAssignmentPayload
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	resourceId := uint(req.ResourceID)

	for _, username := range req.Approvers {
		mh.permissionsRepository.Create(&model.ApprovingPermission{
			ResourceID: resourceId,
			Username:   username,
		})
	}

}

func (mh *ManagementHandlerImpl) RegisterResource(ctx *gin.Context) {
	var req dto.ResourceRegistrationPayload
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	mh.resourcesRepository.Create(&model.Resource{
		SharesCreated:     req.SharesCreated,
		MinSharesRequired: req.MinSharesRequired,
		ResourceDN:        req.ResourceDN,
	})
}
