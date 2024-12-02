// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package mysql

import (
	"github.com/rosas/onex/internal/fakeserver/model"
	"github.com/rosas/onex/internal/fakeserver/store"
	genericstore "github.com/rosas/onex/pkg/store"
	"github.com/rosas/onex/pkg/store/logger/onex"
)

// orders is a concrete implementation of OrderStore that uses a generic store.
type orders struct {
	*genericstore.Store[model.OrderM] // Embedding the generic store for OrderM
}

// Ensure that orders implements the OrderStore interface.
var _ store.OrderStore = (*orders)(nil)

// newOrders creates a new instance of orders.
func newOrders(ds *datastore) *orders {
	return &orders{
		Store: genericstore.NewStore[model.OrderM](ds, onex.NewLogger()),
	}
}
