// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package nightwatch

import (
	store2 "github.com/rosas/onex/internal/gateway/store"
	"github.com/rosas/onex/internal/nightwatch/biz"
	"github.com/rosas/onex/internal/nightwatch/service/v1"
	store4 "github.com/rosas/onex/internal/nightwatch/store"
	"github.com/rosas/onex/internal/nightwatch/validation"
	"github.com/rosas/onex/internal/pkg/client/store"
	store3 "github.com/rosas/onex/internal/usercenter/store"
	"gorm.io/gorm"
)

import (
	_ "github.com/rosas/onex/internal/nightwatch/watcher/all"
)

// Injectors from wire.go:

func wireAggregateStore(db *gorm.DB) (store.Interface, error) {
	datastore := store2.NewStore(db)
	storeDatastore := store3.NewStore(db)
	datastore2 := store.NewStore(datastore, storeDatastore)
	return datastore2, nil
}

func wireService(db *gorm.DB) *v1.NightWatchService {
	datastore := store4.NewStore(db)
	validator := validation.New(datastore)
	bizBiz := biz.NewBiz(datastore)
	nightWatchService := v1.NewNightWatchService(validator, bizBiz)
	return nightWatchService
}

func wireStore(db *gorm.DB) (store4.IStore, error) {
	datastore := store4.NewStore(db)
	return datastore, nil
}