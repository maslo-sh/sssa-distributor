package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"privileges-management/model"
	"privileges-management/model/dto"
	"privileges-management/server/repository"
	"privileges-management/shares"
)

type RequestHandler interface {
	RequestTemporaryAccess(*gin.Context)
	GetAllRequests(*gin.Context)
}

type RequestHandlerImpl struct {
	resourcesRepository   repository.ResourcesRepository
	permissionsRepository repository.ApprovingPermissionsRepository
	requestsRepository    repository.AccessRequestsRepository
}

func NewRequestHandler(resourcesRepo repository.ResourcesRepository, permissionsRepo repository.ApprovingPermissionsRepository, requestsRepo repository.AccessRequestsRepository) RequestHandler {
	return &RequestHandlerImpl{resourcesRepository: resourcesRepo, permissionsRepository: permissionsRepo, requestsRepository: requestsRepo}
}

func (rh *RequestHandlerImpl) RequestTemporaryAccess(ctx *gin.Context) {
	var req dto.RequestAccessPayload
	var resourceId int
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resource := rh.resourcesRepository.Read(uint(req.ResourceID))
	approvingPermissions := rh.permissionsRepository.ReadByResourceId(uint(req.ResourceID))

	var credentials model.Credentials
	credentials, err = shares.GenerateCredentials()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resource.MinSharesRequired > len(approvingPermissions) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("not enough approvers to perform secret sharing"))
		return
	}

	var generatedShares []string
	generatedShares, err = createSecretsFromCredentials(credentials, resource)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	sharesMapping := createSharesMapping(generatedShares, resource, approvingPermissions)

	err = shares.DistributeSecrets(sharesMapping)

	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, err)
		//return
	}

	rh.requestsRepository.Create(&model.AccessRequest{
		Username:        req.Username,
		ResourceID:      uint(resourceId),
		ValidityInHours: req.ValidityInHours,
	})

	ctx.JSON(http.StatusOK, nil)
}

func (rh *RequestHandlerImpl) GetAllRequests(ctx *gin.Context) {
	requests := rh.requestsRepository.ReadAll()
	ctx.JSON(http.StatusOK, requests)
}

func createSharesMapping(shares []string, resource *model.Resource, permissions []model.ApprovingPermission) map[string]string {
	mapping := make(map[string]string)
	for i := 0; float64(i) < math.Min(float64(resource.SharesCreated), float64(resource.MinSharesRequired)); i++ {
		mapping[permissions[i].Username] = shares[i]
	}

	return mapping
}

func createSecretsFromCredentials(credentials model.Credentials, resource *model.Resource) ([]string, error) {
	generatedShares, err := shares.CreateSecretsFromCredentials(resource.MinSharesRequired, resource.SharesCreated, credentials)
	if err != nil {
		return generatedShares, err
	}

	shares.RetrieveCredentialsFromSecrets(generatedShares)

	return generatedShares, nil
}
