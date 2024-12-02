// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package util

import (
	clioptions "github.com/rosas/onex/internal/onexctl/util/options"
	gatewayv1 "github.com/rosas/onex/pkg/api/gateway/v1"
	usercenterv1 "github.com/rosas/onex/pkg/api/usercenter/v1"
)

type Factory interface {
	Login() (token string, err error)
	UserCenterClient() usercenterv1.UserCenterHTTPClient
	GatewayClient() gatewayv1.GatewayHTTPClient
	GetOptions() *clioptions.Options
}
