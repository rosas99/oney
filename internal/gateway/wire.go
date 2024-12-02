// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/rosas/onex/internal/gateway/biz"
	"github.com/rosas/onex/internal/gateway/server"
	"github.com/rosas/onex/internal/gateway/service"
	"github.com/rosas/onex/internal/gateway/store"
	customvalidation "github.com/rosas/onex/internal/gateway/validation"
	"github.com/rosas/onex/internal/pkg/bootstrap"
	"github.com/rosas/onex/internal/pkg/client/usercenter"
	"github.com/rosas/onex/internal/pkg/idempotent"
	"github.com/rosas/onex/internal/pkg/validation"
	"github.com/rosas/onex/pkg/db"
	clientset "github.com/rosas/onex/pkg/generated/clientset/versioned"
	genericoptions "github.com/rosas/onex/pkg/options"
)

// wireApp init kratos application.
func wireApp(
	<-chan struct{},
	bootstrap.AppInfo,
	*server.Config,
	clientset.Interface,
	*db.MySQLOptions,
	*db.RedisOptions,
	*usercenter.UserCenterOptions,
	*genericoptions.RedisOptions,
	*genericoptions.EtcdOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		bootstrap.NewEtcdRegistrar,
		server.ProviderSet,
		store.ProviderSet,
		usercenter.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		validation.ProviderSet,
		idempotent.ProviderSet,
		customvalidation.ProviderSet,
		createInformers,
	)

	return nil, nil, nil
}
