package apis

import (
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs"
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/cfg"
	"github.com/ronannnn/infra/services/jwt"
	"github.com/ronannnn/infra/services/jwt/accesstoken"
	"github.com/ronannnn/infra/services/login"
	"github.com/ronannnn/infra/services/loginrecord"
	"github.com/ronannnn/infra/services/user"
	"go.uber.org/zap"
)

type HttpServer struct {
	infra.BaseHttpServer
	infraMiddleware infra.Middleware
	// services
	loginService       login.Service
	loginRecordService loginrecord.Service
	accessTokenService accesstoken.Service
	jwtService         jwt.Service
	userService        user.Service
	customsSasService  *customs.SasService
}

func NewHttpServer(
	sysCfg *cfg.Sys,
	log *zap.SugaredLogger,
	infraMiddleware infra.Middleware,
	// services
	loginService login.Service,
	loginRecordService loginrecord.Service,
	accessTokenService accesstoken.Service,
	jwtService jwt.Service,
	userService user.Service,
	customsSasService *customs.SasService,
) *HttpServer {
	hs := &HttpServer{
		BaseHttpServer: infra.BaseHttpServer{
			Sys: sysCfg,
			Log: log,
		},
		infraMiddleware: infraMiddleware,
		// services
		loginService:       loginService,
		loginRecordService: loginRecordService,
		accessTokenService: accessTokenService,
		jwtService:         jwtService,
		userService:        userService,
		customsSasService:  customsSasService,
	}
	// golang abstract class reference: https://adrianwit.medium.com/abstract-class-reinvented-with-go-4a7326525034
	hs.BaseHttpServer.HttpServerRunner.HttpServerBaseRunner = hs
	return hs
}
