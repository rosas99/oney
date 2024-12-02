// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package service

import (
	"github.com/rosas/onex/internal/fakeserver/biz"
	v1 "github.com/rosas/onex/pkg/api/fakeserver/v1"
)

type FakeServerService struct {
	v1.UnimplementedFakeServerServer

	biz biz.IBiz
}

func NewFakeServerService(biz biz.IBiz) *FakeServerService {
	return &FakeServerService{biz: biz}
}
