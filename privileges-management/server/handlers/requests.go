package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"privileges-management/model"
	"privileges-management/model/dto"
	"privileges-management/server/repository"
	"privileges-management/sssa"
	"strconv"
)

type RequestHandler interface {
	RequestTemporaryAccess(ctx *gin.Context)
}

type RequestHandlerImpl struct {
	resourcesRepository   repository.ResourcesRepository
	permissionsRepository repository.ApprovingPermissionsRepository
}

func NewRequestHandler(resourcesRepo repository.ResourcesRepository, permissionsRepo repository.ApprovingPermissionsRepository) RequestHandler {
	return &RequestHandlerImpl{resourcesRepository: resourcesRepo, permissionsRepository: permissionsRepo}
}

func (rh *RequestHandlerImpl) RequestTemporaryAccess(ctx *gin.Context) {
	var req dto.RequestAccessPayload
	var resourceId int
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	resourceId, err = strconv.Atoi(req.ResourceID)
	resource := rh.resourcesRepository.Read(uint(resourceId))
	castedResource := resource.(*model.Resource)
	approvingPermissions := rh.permissionsRepository.ReadByResourceId(resourceId)

	var credentials model.Credentials
	credentials, err = sssa.GenerateCredentials()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if castedResource.MinSharesRequired > len(approvingPermissions) {
		ctx.JSON(http.StatusBadRequest, errors.New("not enough approvers to perform secret sharing"))
	}

	var shares []string
	shares, err = createSecretsFromCredentials(credentials, castedResource)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	sharesMapping := createSharesMapping(shares, castedResource, approvingPermissions)

	err = sssa.DistributeSecrets(sharesMapping)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, nil)
}

func createSharesMapping(shares []string, resource *model.Resource, permissions []model.ApprovingPermission) map[string]string {
	mapping := make(map[string]string)
	for i := 0; float64(i) < math.Min(float64(resource.SharesCreated), float64(resource.MinSharesRequired)); i++ {
		mapping[permissions[i].Username] = shares[i]
	}

	return mapping
}

func createSecretsFromCredentials(credentials model.Credentials, resource *model.Resource) ([]string, error) {
	shares, err := sssa.CreateSecretsFromCredentials(resource.MinSharesRequired, resource.SharesCreated, credentials)
	if err != nil {
		return shares, err
	}

	return shares, nil
}
