package server

import (
	"gorm.io/gorm"

	"skadi/backend/config"
	authhttpv1 "skadi/backend/internal/app/auth/controller/http/v1"
	authrepo "skadi/backend/internal/app/auth/repository"
	authuc "skadi/backend/internal/app/auth/usecase"
	classhttpv1 "skadi/backend/internal/app/class/controller/http/v1"
	classrepo "skadi/backend/internal/app/class/repository"
	classuc "skadi/backend/internal/app/class/usecase"
	commenthttpv1 "skadi/backend/internal/app/comment/controller/http/v1"
	commentrepo "skadi/backend/internal/app/comment/repository"
	commentuc "skadi/backend/internal/app/comment/usecase"
	examplehttpv1 "skadi/backend/internal/app/example/controller/http/v1"
	filehttpv1 "skadi/backend/internal/app/file/controller/http/v1"
	filerepo "skadi/backend/internal/app/file/repository"
	fileuc "skadi/backend/internal/app/file/usecase"
	"skadi/backend/internal/app/service/server/middleware"
	solhttpv1 "skadi/backend/internal/app/solution/controller/http/v1"
	solrepo "skadi/backend/internal/app/solution/repository"
	soluc "skadi/backend/internal/app/solution/usecase"
	statusrepo "skadi/backend/internal/app/status/repository"
	taskhttpv1 "skadi/backend/internal/app/task/controller/http/v1"
	taskrepo "skadi/backend/internal/app/task/repository"
	taskuc "skadi/backend/internal/app/task/usecase"
	userhttpv1 "skadi/backend/internal/app/user/controller/http/v1"
	userrepo "skadi/backend/internal/app/user/repository"
	useruc "skadi/backend/internal/app/user/usecase"
	"skadi/backend/internal/pkg/cache"
	"skadi/backend/internal/pkg/validator"
)

// registerEndpointsV1 register all endpoints for 1st version of API.
func (s *Server) registerEndpointsV1(cfg *config.Config, dbStorage *gorm.DB,
	cacheStorage cache.Storage, valid validator.Validator) {

	// create repos
	authRepoCache := authrepo.NewRepoCache(cfg, cacheStorage)
	userRepoDB := userrepo.NewRepoDB(dbStorage)
	classRepoDB := classrepo.NewRepoDB(dbStorage)
	taskRepoDB := taskrepo.NewRepoDB(dbStorage)
	solRepoDB := solrepo.NewRepoDB(dbStorage)
	statusRepoDB := statusrepo.NewRepoDB(dbStorage)
	fileRepoDB := filerepo.NewRepoDB(dbStorage)
	commentRepoDB := commentrepo.NewRepoDB(dbStorage)
	// create usecases
	authUCClient := authuc.NewUCClient(cfg, userRepoDB, authRepoCache)
	authUCMiddleware := authuc.NewUCMiddleware(cfg, authRepoCache)
	userUCAdminClient := useruc.NewUCAdminClient(cfg, userRepoDB, classRepoDB)
	classUCAdminClient := classuc.NewUCAdminClient(cfg, classRepoDB, userRepoDB)
	taskUCTeacher := taskuc.NewUCTeacher(cfg, taskRepoDB, userRepoDB)
	solUCClient := soluc.NewUCClient(cfg, solRepoDB, taskRepoDB)
	solUCStudent := soluc.NewUCStudent(cfg, solRepoDB, statusRepoDB)
	solUCTeacher := soluc.NewUCTeacher(cfg, solRepoDB, statusRepoDB)
	fileUCClient := fileuc.NewUCClient(cfg, fileRepoDB)
	commentUCClient := commentuc.NewUCClient(cfg, commentRepoDB, solRepoDB)
	// create controllers
	authController := authhttpv1.NewController(cfg, authUCClient, valid)
	exampleController := examplehttpv1.NewController()
	userControllerAdmin := userhttpv1.NewControllerAdmin(userUCAdminClient, valid)
	userController := userhttpv1.NewController(userUCAdminClient, valid)
	classControllerAdmin := classhttpv1.NewControllerAdmin(classUCAdminClient, valid)
	classController := classhttpv1.NewController(classUCAdminClient, valid)
	taskControllerTeacher := taskhttpv1.NewControllerTeacher(cfg, taskUCTeacher, valid)
	solController := solhttpv1.NewController(solUCClient, valid)
	solControllerStudent := solhttpv1.NewControllerStudent(cfg, solUCStudent, valid)
	solControllerTeacher := solhttpv1.NewControllerTeacher(solUCTeacher, valid)
	fileController := filehttpv1.NewController(fileUCClient, valid)
	commentController := commenthttpv1.NewController(commentUCClient, valid)

	// middlewares
	mwJWTRefresh := middleware.JWTRefresh(cfg, authUCMiddleware)
	mwJWTAccess := middleware.JWTAccess(cfg)
	// register endpoints
	apiV1 := s.fiberApp.Group("/api/v1")
	if s.Debug() { // register example endpoints if server is in debug mode
		examplehttpv1.RegisterEndpoints(apiV1, exampleController,
			mwJWTAccess, middleware.Allow)
	}
	authhttpv1.RegisterEndpoints(apiV1, authController, mwJWTRefresh)
	userhttpv1.RegisterEndpoints(apiV1, userController, userControllerAdmin,
		mwJWTAccess, middleware.Allow)
	classhttpv1.RegisterEndpoints(apiV1, classController, classControllerAdmin,
		mwJWTAccess, middleware.Allow)
	taskhttpv1.RegisterEndpoints(apiV1, taskControllerTeacher, mwJWTAccess, middleware.Allow)
	solhttpv1.RegisterEndpoints(apiV1, solController, solControllerStudent, solControllerTeacher,
		mwJWTAccess, middleware.Allow)
	filehttpv1.RegisterEndpoints(apiV1, fileController, mwJWTAccess, middleware.Allow)
	commenthttpv1.RegisterEndpoints(apiV1, commentController, mwJWTAccess, middleware.Allow)
}
