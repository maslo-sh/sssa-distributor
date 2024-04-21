package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v3"
	"gorm.io/gorm"
	"privileges-management/database"
	"privileges-management/server/handlers"
	"privileges-management/server/repository"
)

func NewRouter() *gin.Engine {
	service := gin.Default()

	db, ldapConn, err := createConnectors()

	if err != nil {
		panic(err)
	}

	permissionsRepository, resourcesRepository, requestsRepository := createRepositories(db)

	reqHandler, approveHandler, managementHandler, loginHandler := createHandlers(permissionsRepository, resourcesRepository, requestsRepository, ldapConn)

	router := service.Group("/api")
	loginRouter := router.Group("/login")
	loginRouter.POST("", loginHandler.Login)

	requestRouter := router.Group("/request")
	requestRouter.GET("", reqHandler.GetAllRequests)
	requestRouter.PUT("", reqHandler.RequestTemporaryAccess)

	managementRouter := router.Group("/management")
	managementRouter.PUT("/approver", managementHandler.AssignApproversToResource)
	managementRouter.POST("/resource", managementHandler.RegisterResource)
	managementRouter.POST("/approver", managementHandler.RegisterApprover)

	approveRouter := router.Group("/approvals")
	approveRouter.PUT("/deny", approveHandler.Deny)
	approveRouter.PUT("/approve", approveHandler.Approve)

	return service
}

func createRepositories(db *gorm.DB) (
	repository.ApprovingPermissionsRepository,
	repository.ResourcesRepository,
	repository.AccessRequestsRepository) {
	return repository.NewApprovingPermissionsRepository(db), repository.NewResourcesRepository(db), repository.NewAccessRequestRepository(db)
}

func createHandlers(
	approvingPermsRepo repository.ApprovingPermissionsRepository,
	resourcesRepo repository.ResourcesRepository,
	accessRequestRepo repository.AccessRequestsRepository,
	conn *ldap.Conn) (handlers.RequestHandler, handlers.ApprovalsHandler, handlers.ManagementHandler, handlers.LoginHandler) {

	return handlers.NewRequestHandler(resourcesRepo, approvingPermsRepo, accessRequestRepo), handlers.NewApprovalsHandler(accessRequestRepo),
		handlers.NewManagementHandler(approvingPermsRepo, resourcesRepo), handlers.NewLoginHandler(conn)
}

func createConnectors() (db *gorm.DB, conn *ldap.Conn, err error) {
	ldapConn, err := database.ConnectToAD()
	if err != nil {
		//return nil, nil, err
	}

	sqlDb, err := database.ConnectToDatabase()
	if err != nil {
		return nil, nil, err
	}

	return sqlDb, ldapConn, nil
}
