package apis

import (
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs"
	"github.com/ronannnn/infra/cfg"
	"github.com/ronannnn/infra/handler"
	"github.com/ronannnn/infra/services/apirecord"
	"github.com/ronannnn/infra/services/jwt"
	"github.com/ronannnn/infra/services/jwt/accesstoken"
	"github.com/ronannnn/infra/services/login"
	"github.com/ronannnn/infra/services/loginrecord"
	"github.com/ronannnn/infra/services/user"
	"go.uber.org/zap"
)

type HttpServer struct {
	handler.BaseHttpServer
	h handler.HttpHandler
	// middlewares
	handlerMw     handler.Middleware
	accessTokenMw accesstoken.Middleware
	apiRecordMw   apirecord.Middleware
	// services
	loginService       login.Service
	loginRecordService loginrecord.Service
	accessTokenService accesstoken.Service
	jwtService         jwt.Service
	userService        user.Service
	customsService     customs.CustomsService
	customsSasService  *customs.SasService
	customsDecService  *customs.DecService
}

func NewHttpServer(
	sysCfg *cfg.Sys,
	log *zap.SugaredLogger,
	h handler.HttpHandler,
	// middlewares
	handlerMw handler.Middleware,
	accessTokenMw accesstoken.Middleware,
	apiRecordMw apirecord.Middleware,
	// services
	loginService login.Service,
	loginRecordService loginrecord.Service,
	accessTokenService accesstoken.Service,
	jwtService jwt.Service,
	userService user.Service,
	customsService customs.CustomsService,
	customsSasService *customs.SasService,
	customsDecService *customs.DecService,
) *HttpServer {
	hs := &HttpServer{
		BaseHttpServer: handler.BaseHttpServer{
			Sys: sysCfg,
			Log: log,
		},
		h: h,
		// middlewares
		handlerMw:     handlerMw,
		accessTokenMw: accessTokenMw,
		apiRecordMw:   apiRecordMw,
		// services
		loginService:       loginService,
		loginRecordService: loginRecordService,
		accessTokenService: accessTokenService,
		jwtService:         jwtService,
		userService:        userService,
		customsService:     customsService,
		customsSasService:  customsSasService,
		customsDecService:  customsDecService,
	}
	// golang abstract class reference: https://adrianwit.medium.com/abstract-class-reinvented-with-go-4a7326525034
	hs.BaseHttpServer.HttpServerRunner.HttpServerBaseRunner = hs
	customsService.ListenImpPath()
	return hs
}
