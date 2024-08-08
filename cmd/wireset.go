package main

import (
	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/apis"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/sas"
	"github.com/ronannnn/gx-customs-bridge/internal/services/db"

	"github.com/google/wire"
	"github.com/ronannnn/infra"
	"github.com/ronannnn/infra/services/apirecord"
	"github.com/ronannnn/infra/services/jwt"
	"github.com/ronannnn/infra/services/jwt/accesstoken"
	"github.com/ronannnn/infra/services/jwt/refreshtoken"
	"github.com/ronannnn/infra/services/login"
	"github.com/ronannnn/infra/services/loginrecord"
	"github.com/ronannnn/infra/services/user"
)

var wireSet = wire.NewSet(
	// configs
	internal.ProvideCfg,
	internal.ProvideSysCfg,
	internal.ProvideLogCfg,
	internal.ProvideDbCfg,
	internal.ProvideAuthCfg,
	internal.ProvideUserCfg,
	internal.ProvideCustomsCfg,
	// infra
	infra.ProvideCasbinEnforcer,
	db.ProvideService,
	internal.ProvideLog,
	// middleware
	infra.ProvideMiddleware,
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
	customs.ProvideSasService,
	customs.ProvideCustomsService,
	// stores
	refreshtoken.ProvideStore,
	user.ProvideStore,
	loginrecord.ProvideStore,
	// server
	apis.NewHttpServer,
)
