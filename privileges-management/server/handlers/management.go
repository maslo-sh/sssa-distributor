package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"privileges-management/model/dto"
	"privileges-management/server/repository"
)

type ManagementHandler interface {
	AssignApproversToResource(*gin.Context)
}

type ManagementHandlerImpl struct {
	permissionsRepository repository.ApprovingPermissionsRepository
}

func NewManagementHandler(permissionsRepo repository.ApprovingPermissionsRepository) ManagementHandler {
	return &ManagementHandlerImpl{permissionsRepository: permissionsRepo}
}

func (mh *ManagementHandlerImpl) AssignApproversToResource(ctx *gin.Context) {
	var req dto.ApproverAssignmentPayload
	//var resourceId int
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	//resourceId, err = strconv.Atoi(req.ResourceID)
	//resource := rh.resourcesRepository.Read(uint(resourceId))
	//castedResource := resource.(*model.Resource)

}
