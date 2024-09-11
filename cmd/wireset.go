package main

import (
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/apis"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/dec"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/ronannnn/gx-customs-bridge/internal/services/db"
	"github.com/ronannnn/gx-customs-bridge/internal/services/rmq"

	"github.com/google/wire"
	"github.com/ronannnn/infra/handler"
	"github.com/ronannnn/infra/i18n"
	"github.com/ronannnn/infra/services/apirecord"
	"github.com/ronannnn/infra/services/jwt"
	"github.com/ronannnn/infra/services/jwt/accesstoken"
	"github.com/ronannnn/infra/services/jwt/refreshtoken"
	"github.com/ronannnn/infra/services/login"
	"github.com/ronannnn/infra/services/loginrecord"
	"github.com/ronannnn/infra/services/user"
	"github.com/ronannnn/infra/validator"
)

var wireSet = wire.NewSet(
	// configs
	internal.ProvideCfg,
	internal.ProvideUserCfg,
	internal.ProvideSysCfg,
	internal.ProvideLogCfg,
	internal.ProvideDbCfg,
	internal.ProvideAccessTokenCfg,
	internal.ProvideRefreshTokenCfg,
	internal.ProvideRabbitmqCfg,
	internal.ProvideI18nCfg,
	internal.ProvideCustomsCfg,
	// infra
	db.ProvideService,
	internal.ProvideLog,
	rmq.ProvideService,
	i18n.New,
	validator.New,
	handler.NewHttpHandler,
	// middleware
	handler.ProvideMiddleware,
	accesstoken.ProvideMiddleware,
	apirecord.ProvideMiddleware,
	// services
	jwt.ProvideService,
	accesstoken.ProvideService,
	refreshtoken.ProvideService,
	login.ProvideService,
	loginrecord.ProvideService,
	user.ProvideService,
	apirecord.ProvideService,
	apirecord.ProvideStore,
	sas.ProvideSasXmlService,
	dec.ProvideDecXmlService,
	customs.ProvideSasService,
	customs.ProvideDecService,
	customs.ProvideCustomsService,
	common.ProvideCustomsCommonXmlService,
	// stores
	refreshtoken.ProvideStore,
	user.ProvideStore,
	loginrecord.ProvideStore,
	// server
	apis.NewHttpServer,
)
